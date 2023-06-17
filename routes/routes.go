package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harold/attendance-golang/res/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", handlers.CreateUser)
	api.Post("/login", handlers.Login)
	api.Get("/get_users/:id", handlers.GetUserByID)
	api.Get("/users", handlers.GetUsers)
}
