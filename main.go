package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/email_notification_server/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	handlers.Routes(app)

	log.Fatal(app.Listen(":3000"))

}
