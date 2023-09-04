package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/schoolservice"

	userDomain "github.com/k61b/okul/internal/domain/user"
)

type SchoolHandlers struct {
	schoolService *schoolservice.SchoolService
}

func NewSchoolHandlers(schoolService *schoolservice.SchoolService) *SchoolHandlers {
	return &SchoolHandlers{schoolService: schoolService}
}

// @Summary Create a school
// @Description Create a school
// @Tags schools
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param description body string true "Description"
// @Param address body string true "Address"
// @Param phone_number body string true "Phone Number"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /schools [post]
func (h *SchoolHandlers) CreateSchoolHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	owner_email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err = h.schoolService.CreateSchool(
		body.Name,
		body.Description,
		body.Address,
		body.PhoneNumber,
		owner_email,
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "School created successfully",
	})
}

// @Summary Get all schools
// @Description Get all schools
// @Tags schools
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /schools [get]
func (h *SchoolHandlers) GetAllSchoolsHandler(c *fiber.Ctx) error {
	schools, err := h.schoolService.GetAllSchools()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"schools": schools,
	})
}

// @Summary Get school by ID
// @Description Get school by ID
// @Tags schools
// @Accept json
// @Produce json
// @Param id path int true "School ID"
// @Success 200 {object} string "OK"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /schools/{id} [get]
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

// @Summary Update a school
// @Description Update a school
// @Tags schools
// @Accept json
// @Produce json
// @Param id path int true "School ID"
// @Param name body string true "Name"
// @Param description body string true "Description"
// @Param address body string true "Address"
// @Param phone_number body string true "Phone Number"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /schools/{id} [put]
func (h *SchoolHandlers) UpdateSchoolHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")
	params := c.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	school, err := h.schoolService.GetSchoolByID(id)
	if err != nil {
		return err
	}

	if school.OwnerEmail != email {
		return fiber.NewError(fiber.StatusForbidden, "You are not authorized to update this school")
	}

	school.Name = body.Name
	school.Description = body.Description
	school.Address = body.Address
	school.PhoneNumber = body.PhoneNumber
	school.UpdatedAt = time.Now()

	err = h.schoolService.UpdateSchool(school)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "School updated successfully",
	})
}

// @Summary Delete a school
// @Description Delete a school
// @Tags schools
// @Accept json
// @Produce json
// @Param id path int true "School ID"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /schools/{id} [delete]
func (h *SchoolHandlers) SuspendSchoolHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")
	params := c.Params("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		return err
	}

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	school, err := h.schoolService.GetSchoolByID(id)
	if err != nil {
		return err
	}

	if school.OwnerEmail != email {
		return fiber.NewError(fiber.StatusForbidden, "You are not authorized to suspend this school")
	}

	err = h.schoolService.SuspendSchool(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "School suspended successfully",
	})
}
