package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/userservice"
	"github.com/k61b/okul/internal/application/verificationservice"

	userDomain "github.com/k61b/okul/internal/domain/user"
	verificationDomain "github.com/k61b/okul/internal/domain/verification"
)

type UserHandlers struct {
	userService         *userservice.UserService
	verificationService *verificationservice.VerificationService
}

func NewUserHandlers(userService *userservice.UserService, verificationService *verificationservice.VerificationService) *UserHandlers {
	return &UserHandlers{userService: userService, verificationService: verificationService}
}

func (h *UserHandlers) SessionHandler(c *fiber.Ctx) error {
	var u userDomain.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(u.Email)
	if err != nil {
		return err
	}

	if user == nil {
		hash, err := userDomain.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash

		user, err = h.userService.CreateUser(u.Email, u.Password, u.Name, u.Surname)
		if err != nil {
			return err
		}
	} else {
		if !userDomain.CheckPassword(u.Password, user.Password) {
			return fiber.ErrUnauthorized
		}
	}

	token, err := userDomain.GenerateJWTToken(user.Email)
	if err != nil {
		return err
	}

	c.Cookie(userDomain.GenerateCookie(token))

	return c.JSON(fiber.Map{"token": token})
}

func (h *UserHandlers) LogoutHandler(c *fiber.Ctx) error {
	c.Cookie(userDomain.GenerateCookie(""))
	return c.JSON(fiber.Map{"message": "success"})
}

func (h *UserHandlers) MeHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	email, err := userDomain.GetEmailFromToken(token)
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

	var u userDomain.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		return err
	}

	u.ID = user.ID
	u.Email = user.Email
	u.IsEmailVerified = user.IsEmailVerified
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

func (h *UserHandlers) SendVerificationEmailAndStoreHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")
	verificationType := "email"

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(time.Hour * 24)

	verificationToken, err := verificationDomain.GenerateVerificationToken(email, expiresAt)
	if err != nil {
		return err
	}

	if err := verificationDomain.SendVerificationEmail(email, verificationToken); err != nil {
		return err
	}

	if err := h.verificationService.CreateVerification(verificationType, email, verificationToken, expiresAt); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}

func (h *UserHandlers) VerifyEmailHandler(c *fiber.Ctx) error {
	token := c.Query("token")

	id, verificationType, email, err := h.verificationService.GetVerificationInfoFromToken(token)
	if err != nil {
		return err
	}

	if verificationType != "email" && email == "" {
		return fiber.ErrBadRequest
	}

	if err := h.userService.UpdateUserEmailVerificationStatus(email, true); err != nil {
		return err
	}

	if err := h.verificationService.DeleteVerification(id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}

func (h *UserHandlers) ForgotPasswordHandler(c *fiber.Ctx) error {
	var u userDomain.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(u.Email)
	if err != nil {
		return err
	}

	if user == nil {
		return fiber.ErrNotFound
	}

	expiresAt := time.Now().Add(time.Hour * 24)

	verificationToken, err := verificationDomain.GenerateVerificationToken(u.Email, expiresAt)
	if err != nil {
		return err
	}

	if err := verificationDomain.SendVerificationEmailForPassword(u.Email, verificationToken); err != nil {
		return err
	}

	if err := h.verificationService.CreateVerification("password", u.Email, verificationToken, expiresAt); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}

func (h *UserHandlers) ResetPasswordHandler(c *fiber.Ctx) error {
	var u userDomain.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	token := c.Query("token")

	id, verificationType, email, err := h.verificationService.GetVerificationInfoFromToken(token)
	if err != nil {
		return err
	}

	if verificationType != "password" && email == "" {
		return fiber.ErrBadRequest
	}

	hash, err := userDomain.HashPassword(u.Password)
	if err != nil {
		return err
	}

	if err := h.userService.UpdateUserPassword(email, hash); err != nil {
		return err
	}

	if err := h.verificationService.DeleteVerification(id); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "success"})
}
