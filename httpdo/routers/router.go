package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/self-denying/loginlimiter/httpdo/action"
	"log"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2022/6/21
* Time: 15:44
**/

func Run() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:username", action.Join)
	app.Post("/send", action.Send)

	log.Fatal(app.Listen(":3000"))
}
