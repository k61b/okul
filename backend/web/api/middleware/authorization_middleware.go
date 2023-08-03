package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement your authorization logic here
		// For example, check if the user has the required role to access the route
		// If authorization fails, return an appropriate response (e.g., 403 Forbidden)

		// If authorization succeeds, proceed to the next middleware or handler
		return c.Next()
	}
}
