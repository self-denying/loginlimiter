package handles

import (
	"errors"
	"github.com/self-denying/loginlimiter/kernel/constant"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 11:25
**/

type Message struct {
	Type      constant.MessageType `json:"type"`
	Content   string               `json:"content,omitempty"`
	UserName  string               `json:"username,omitempty"`
	NewClient *Client              `json:"-"`
	Extra     interface{}          `json:"extra,omitempty"`
}

func (m *Message) TODO(w *WsManager) error {
	if m.Type <= 0 {
		return errors.New("该类型消息暂不支持")
	}
	switch m.Type {
	case constant.KickOut:
		return w.KickOut(m.UserName)
	case constant.BroadCast:
		return w.BroadCast([]byte(m.Content))
	case constant.OnlyOne:
		return w.OnlyOne(m.UserName, []byte(m.Content))
	}
	return nil
}
