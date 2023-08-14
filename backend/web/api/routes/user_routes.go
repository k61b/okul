package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupUserRoutes(app *fiber.App, userHandlers *handlers.UserHandlers) {
	user := app.Group("/user")

	user.Use(middleware.AuthMiddleware)

	user.Post("/", userHandlers.SessionHandler)
	user.Post("/logout", userHandlers.LogoutHandler)
}
