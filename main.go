package main

import (
	"github.com/self-denying/loginlimiter/httpdo/routers"
	"github.com/self-denying/loginlimiter/kernel"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 14:46
**/

func main() {

	go kernel.WebsocketConnPool.Listener()

	routers.Run()
}
