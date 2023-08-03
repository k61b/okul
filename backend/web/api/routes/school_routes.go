// web/api/routes/school_routes.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupSchoolRoutes(app *fiber.App, schoolHandlers *handlers.SchoolHandlers) {
	school := app.Group("/school")

	school.Use(middleware.AuthMiddleware)
	// Add any other middleware specific to schools

	school.Post("/", schoolHandlers.CreateSchoolHandler)
	// Define more school routes here
}
