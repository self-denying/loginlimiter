package action

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/self-denying/loginlimiter/kernel"
	"github.com/self-denying/loginlimiter/kernel/handles"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 15:09
**/

func Join(ctx *fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		username := c.Params("username")
		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		client := &handles.Client{
			Conn: c,
		}
		_ = kernel.WebsocketConnPool.Join(username, client)
		client.Listener()
	})(ctx)
}

func Send(ctx *fiber.Ctx) error {
	scanner := new(handles.Message)
	err := ctx.BodyParser(scanner)
	if err != nil {
		return ctx.JSON(fmt.Sprintf("请求参数错误:%v", err.Error()))
	}
	kernel.WebsocketConnPool.Message <- scanner
	return ctx.JSON("处理成功")
}
