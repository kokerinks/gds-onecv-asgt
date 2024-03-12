package routes

import "github.com/gofiber/fiber/v2"

func ApiHandling(router fiber.Router) {
	router.Get("/register", func(c *fiber.Ctx) error {
		return c.SendString("Register")
	})
	router.Get("/commonstudents", func(c *fiber.Ctx) error {
		return c.SendString("Common Students")
	})
	router.Post("/suspend", func(c *fiber.Ctx) error {
		return c.SendString("Suspend")
	})
	router.Post("/retrievefornotifications", func(c *fiber.Ctx) error {
		return c.SendString("Retrieve for Notifications")
	})
}
