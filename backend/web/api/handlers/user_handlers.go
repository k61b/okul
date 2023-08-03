package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/userservice"
)

type UserHandlers struct {
	userService *userservice.UserService
}

func NewUserHandlers(userService *userservice.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

// Implement HTTP request handlers for user-related actions here
// For example: CreateUserHandler, GetUserByIDHandler, UpdateUserHandler, DeleteUserHandler, etc.
func (h *UserHandlers) CreateUserHandler(c *fiber.Ctx) error {
	// Implement create user logic and return appropriate response
	return nil
}
