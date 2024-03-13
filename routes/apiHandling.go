package routes

import (
	"gds-onecv-asgt/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApiHandling(router fiber.Router) {
	router.Post("/register", controllers.Register)
	router.Get("/commonstudents", controllers.CommonStudents)
	router.Post("/suspend", controllers.Suspend)
	router.Post("/retrievefornotifications", controllers.RetrieveForNotifications)
}
