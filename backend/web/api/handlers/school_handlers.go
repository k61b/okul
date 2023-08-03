package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/schoolservice"
)

type SchoolHandlers struct {
	schoolService *schoolservice.SchoolService
}

func NewSchoolHandlers(schoolService *schoolservice.SchoolService) *SchoolHandlers {
	return &SchoolHandlers{schoolService: schoolService}
}

// Implement HTTP request handlers for school-related actions here
// For example: CreateSchoolHandler, GetSchoolByIDHandler, UpdateSchoolHandler, etc.
func (h *SchoolHandlers) CreateSchoolHandler(c *fiber.Ctx) error {
	// Implement create school logic and return appropriate response
	return nil
}
