package handles

import (
	"github.com/gofiber/websocket/v2"
	"log"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 10:19
**/

type Client struct {
	Conn *websocket.Conn
}

func (c *Client) Listener() {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.Conn.ReadMessage(); err != nil {
			_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.Conn.WriteMessage(mt, msg); err != nil {
			_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			break
		}
	}
}
