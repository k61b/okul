package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/adminservice"
)

type AdminHandlers struct {
	adminService *adminservice.AdminService
}

func NewAdminHandlers(adminService *adminservice.AdminService) *AdminHandlers {
	return &AdminHandlers{adminService: adminService}
}

// Implement HTTP request handlers for admin-related actions here
// For example: CreateAdminHandler, GetAdminByIDHandler, DeleteAdminHandler, etc.
func (h *AdminHandlers) CreateAdminHandler(c *fiber.Ctx) error {
	// Implement create admin logic and return appropriate response
	return nil
}
