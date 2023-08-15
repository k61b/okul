package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
)

func SetupUserRoutes(app *fiber.App, userHandlers *handlers.UserHandlers) {
	user := app.Group("/user")

	user.Post("/", userHandlers.SessionHandler)
	user.Post("/logout", userHandlers.LogoutHandler)
}
