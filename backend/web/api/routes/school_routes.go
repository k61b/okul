// web/api/routes/school_routes.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
)

func SetupSchoolRoutes(app *fiber.App, schoolHandlers *handlers.SchoolHandlers) {
	school := app.Group("/school")

	school.Post("/", schoolHandlers.CreateSchoolHandler)
	school.Get("/", schoolHandlers.GetAllSchoolsHandler)
	school.Get("/:id", schoolHandlers.GetSchoolByIDHandler)
}
