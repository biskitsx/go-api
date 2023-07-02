package routes

import (
	"github.com/biskitsx/go-api/controllers"
	"github.com/biskitsx/go-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// auth
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)

	// post
	app.Get("/api/posts/", controllers.GetPosts)
	app.Get("/api/posts/:id", controllers.GetPost)
	app.Post("/api/posts", middleware.VerifyUser, controllers.CreatePost)
}
