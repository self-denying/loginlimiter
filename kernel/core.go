package kernel

import (
	"github.com/self-denying/loginlimiter/kernel/handles"
	"sync"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 14:55
**/

var WebsocketConnPool = &handles.WsManager{
	Clients: make(map[string]*handles.Client),
	Message: make(chan *handles.Message, 100),
	Mutex:   sync.Mutex{},
}
