package main

import (
	"github.com/biskitsx/go-api/db"
	"github.com/biskitsx/go-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.Connect()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	routes.SetupRoutes(app)
	app.Listen(":8080")
}
