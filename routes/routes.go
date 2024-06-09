package routes

import (
	AuthController "github.com/go-auth/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// define a group
	api := app.Group("/api")

	// auth routes
	api.Post("/register", AuthController.Register)
	api.Post("/login", AuthController.Login)
}