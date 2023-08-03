package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupAdminRoutes(app *fiber.App, adminHandlers *handlers.AdminHandlers) {
	admin := app.Group("/admin")

	admin.Use(middleware.AuthMiddleware)
	admin.Use(middleware.AuthorizationMiddleware("admin")) // Requires "admin" role

	admin.Post("/", adminHandlers.CreateAdminHandler)
	// Define more admin routes here
}
