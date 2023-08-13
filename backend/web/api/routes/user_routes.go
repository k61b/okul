package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupUserRoutes(app *fiber.App, userHandlers *handlers.UserHandlers) {
	user := app.Group("/user")

	user.Use(middleware.AuthMiddleware)
	// Add any other middleware specific to users

	user.Post("/", userHandlers.SessionHandler)
	// Define more user routes here
}
