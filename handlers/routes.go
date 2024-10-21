package handlers

import "github.com/gofiber/fiber/v2"

func Routes(routes *fiber.App) {

	routes.Post("/send_email", SendEmail)
}
