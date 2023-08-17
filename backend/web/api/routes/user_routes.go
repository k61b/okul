package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupUserRoutes(app *fiber.App, userHandlers *handlers.UserHandlers) {
	user := app.Group("/user")

	user.Post("/", userHandlers.SessionHandler)
	user.Post("/logout", userHandlers.LogoutHandler)

	user.Get("/me", middleware.AuthMiddleware, userHandlers.MeHandler)

	user.Put("/:id", middleware.AuthMiddleware, userHandlers.UpdateHandler)
	user.Delete("/:id", middleware.AuthMiddleware, userHandlers.DeleteHandler)
}
