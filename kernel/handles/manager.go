package handles

import (
	"errors"
	"fmt"
	"github.com/gofiber/websocket/v2"
	"github.com/self-denying/loginlimiter/kernel/constant"
	"log"
	"sync"
	"time"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 9:49
**/

type WsManager struct {
	Clients map[string]*Client
	Mutex   sync.Mutex
	Message chan *Message //json格式
}

//Add a user to the websocket queue

func (w *WsManager) Join(username string, profile *Client) error {
	if profile == nil || username == "" {
		return errors.New("非法用户")
	}
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	w.Clients[username] = profile
	return nil
}

//Kick a user off the websocket queue

func (w *WsManager) KickOut(username string) error {
	if username == "" {
		return errors.New("无效的username")
	}
	w.Mutex.Lock()
	defer w.Mutex.Unlock()
	delete(w.Clients, username)
	return nil
}

//Broadcast message to all users

func (w *WsManager) BroadCast(message []byte) error {
	for k, v := range w.Clients {
		err := v.Conn.WriteMessage(1, message)
		if err != nil {
			return fmt.Errorf("发送消息给%v用户异常:%v", k, err.Error())
		}
	}
	return nil
}

//Send a message to one person alone

func (w *WsManager) OnlyOne(username string, message []byte) error {
	client, ok := w.Clients[username]
	if !ok {
		return fmt.Errorf("无%v用户", username)
	}
	return client.Conn.WriteMessage(1, message)
}

func (w *WsManager) Listener() {
	for {
		select {
		case msg := <-w.Message:
			err := msg.TODO(w)
			if err != nil {
				log.Println("WsManagerTODOErr:", err.Error(), fmt.Sprintf("%#v", *msg))
			}
		case <-time.Tick(time.Second * 3):
			for k, v := range w.Clients {
				if v.Conn == nil || v.Conn.Conn == nil {
					log.Println(fmt.Sprintf("%v客户端已断开,即将删除该连接", k))
					_ = w.KickOut(k)
					continue
				}
				_ = v.Conn.SetWriteDeadline(time.Now().Add(constant.WriteWait))
				if err := v.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Println(fmt.Sprintf("%v客户端已断开,即将删除该连接", k))
					_ = w.KickOut(k)
				}
			}
		}
	}
}
