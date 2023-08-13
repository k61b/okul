package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/userservice"
	domain "github.com/k61b/okul/internal/domain/user"
)

type UserHandlers struct {
	userService *userservice.UserService
}

func NewUserHandlers(userService *userservice.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

func (h *UserHandlers) SessionHandler(c *fiber.Ctx) error {
	var u domain.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(u.Email)
	if err != nil {
		return err
	}

	if user == nil {
		hash, err := domain.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash

		user, err = h.userService.CreateUser(u.Email, u.Password, u.Name, u.Surname)
		if err != nil {
			return err
		}
	} else {
		if !domain.CheckPassword(u.Password, user.Password) {
			return fiber.ErrUnauthorized
		}
	}

	token, err := domain.GenerateJWTToken(user.Email)
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"token": token})
}
