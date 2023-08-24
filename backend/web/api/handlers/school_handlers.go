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

func (h *SchoolHandlers) CreateSchoolHandler(c *fiber.Ctx) error {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
		OwnerID     int    `json:"owner_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := h.schoolService.CreateSchool(
		body.Name,
		body.Description,
		body.Address,
		body.PhoneNumber,
		body.OwnerID,
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "School created successfully",
	})
}

func (h *SchoolHandlers) GetAllSchoolsHandler(c *fiber.Ctx) error {
	schools, err := h.schoolService.GetAllSchools()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"schools": schools,
	})
}

func (h *SchoolHandlers) GetSchoolByIDHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	school, err := h.schoolService.GetSchoolByID(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"school": school,
	})
}
