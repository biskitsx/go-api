package main

import (
	"github.com/biskitsx/go-api/db"
	"github.com/biskitsx/go-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load()
	// database connection
	db.Connect()
	// fiber app
	app := fiber.New()
	// routes
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
