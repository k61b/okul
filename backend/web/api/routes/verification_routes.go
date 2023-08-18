package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupVerficationRoutes(app *fiber.App, verificationHandlers *handlers.VerificationHandler) {
	verification := app.Group("/verification")

	verification.Post("/verify-email", middleware.AuthMiddleware, verificationHandlers.VerifyEmailHandler)

}
