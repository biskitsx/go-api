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
	app.Patch("/api/posts/:id", middleware.VerifyUser, controllers.UpdatePost)
	app.Delete("/api/posts/:id", middleware.VerifyUser, controllers.DeletePost)

	// user
	app.Get("/api/users/", controllers.GetUsers)
	app.Get("/api/users/:id", controllers.GetUser)
}
