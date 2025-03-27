package router

import (
	noteRoutes "example.org/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Note Routes
	noteRoutes.SetupNoteRoutes(api)
}
