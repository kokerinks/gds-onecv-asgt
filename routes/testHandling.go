package routes

import "github.com/gofiber/fiber/v2"

func TestHandling(app *fiber.App) {
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
