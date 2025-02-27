package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Listen on PORT 3000
	app.Listen(":3000")
}
