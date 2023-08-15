package middleware

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/k61b/okul/internal/domain/user"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")

	if _, err := domain.ParseToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
