package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gds-onecv-asgt/utils"
	"gds-onecv-asgt/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "govtech-opencv",
	})

	app.Use(cors.New())
	utils.ConnectToDB()
	
	routes.TestHandling(app)
	log.Fatal(app.Listen(":3000"))
}
