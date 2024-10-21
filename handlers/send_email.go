package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5603/email_notification_server/models"
	"gopkg.in/mail.v2"
)

func SendEmail(ctx *fiber.Ctx) error {
	var payload models.EmailRequest

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "false", "error": err.Error()})
	}

	from := "shemanth.kgp@gmail.com"
	password := "dydbmplsdbljitem"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	m := mail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", payload.To)
	m.SetHeader("Subject", "Weather Alert Notification")
	m.SetBody("text/html", fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; text-align: center;">
				<h2>Temperature Alert Notification</h2>
				<p>The temperature has reached the alert threshold in your area.</p>
				<div style="margin: 20px auto; padding: 10px 20px; background-color: #FF5733; color: white; font-size: 24px; font-weight: bold; display: inline-block; border-radius: 5px;">
					Temperature: %s
				</div>
				<p style="font-size: 18px; font-weight: bold;">Weather Condition: %s</p>
				<p style="font-size: 18px; font-weight: bold;">Location: %s</p>
				<p style="font-size: 18px; font-weight: bold;">Date: %s</p>
				<p style="font-size: 18px; font-weight: bold;">Time: %s</p>
				<br/>
				<p style="font-size: 16px;">Stay alert and take precautions accordingly.</p>
			</body>
		</html>`, payload.Temperature, payload.Weather, payload.Location, payload.Date, payload.Time))

	d := mail.NewDialer(smtpHost, smtpPort, from, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).
		JSON(fiber.Map{"status": "true", "message": "User Notified Successfully"})
}
