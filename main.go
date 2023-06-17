package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harold/attendance-golang/routes"
)

func main() {
	app := fiber.New()

	// Register routes and handlers
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
