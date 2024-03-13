package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gds-onecv-asgt/routes"
	"gds-onecv-asgt/testData"
	"gds-onecv-asgt/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "govtech-opencv",
	})

	app.Use(cors.New())
	utils.ConnectToDB(false)
	testData.SeedData()

	apiGroup := app.Group("/api")
	routes.ApiHandling(apiGroup)

	log.Fatal(app.Listen(":3000"))
}
