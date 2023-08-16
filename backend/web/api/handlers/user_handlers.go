package handlers

import (
	"strconv"

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

	c.Cookie(domain.GenerateCookie(token))

	return c.JSON(fiber.Map{"token": token})
}

func (h *UserHandlers) LogoutHandler(c *fiber.Ctx) error {
	c.Cookie(domain.GenerateCookie(""))
	return c.JSON(fiber.Map{"message": "success"})
}

func (h *UserHandlers) MeHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	email, err := domain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(email)
	if err != nil {
		return err
	}

	user.Password = "***"

	return c.JSON(user)
}

func (h *UserHandlers) UpdateHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var u domain.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		return err
	}

	u.ID = user.ID
	u.Email = user.Email
	u.Password = user.Password

	updatedUser, err := h.userService.Update(&u)
	if err != nil {
		return err
	}

	return c.JSON(updatedUser)
}

func (h *UserHandlers) DeleteHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err := h.userService.Delete(id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}
