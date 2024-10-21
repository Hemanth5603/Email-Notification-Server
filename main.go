package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hemanth5603/email_notification_server/handlers"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // You can specify specific origins here, e.g., "https://your-vercel-app.vercel.app"
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	handlers.Routes(app)

	log.Fatal(app.Listen(":3000"))

}
