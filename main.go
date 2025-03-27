package main

import (
	"github.com/gofiber/fiber/v2"

	"example.org/database"
	"example.org/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the database
	database.ConnectDB()

	// Setup the routes
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
