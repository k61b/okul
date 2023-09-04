package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/internal/application/favoriteservice"
	"github.com/k61b/okul/internal/application/schoolservice"
	"github.com/k61b/okul/internal/application/userservice"

	userDomain "github.com/k61b/okul/internal/domain/user"
)

type FavoriteHandlers struct {
	favoriteService *favoriteservice.FavoriteService
	schoolService   *schoolservice.SchoolService
	userService     *userservice.UserService
}

func NewFavoriteHandlers(favoriteService *favoriteservice.FavoriteService, schoolService *schoolservice.SchoolService, userService *userservice.UserService) *FavoriteHandlers {
	return &FavoriteHandlers{favoriteService: favoriteService, schoolService: schoolService, userService: userService}
}

// @Summary Create a favorite
// @Description Create a favorite
// @Tags favorites
// @Accept json
// @Produce json
// @Param schoolID path int true "School ID"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /favorites/{schoolID} [post]
func (h *FavoriteHandlers) CreateFavoriteHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(email)
	if err != nil {
		return err
	}

	schoolID, err := strconv.Atoi(c.Params("schoolID"))
	if err != nil {
		return err
	}

	err = h.favoriteService.CreateFavorite(user.ID, schoolID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @Summary Get favorite schools
// @Description Get favorite schools
// @Tags favorites
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /favorites [get]
func (h *FavoriteHandlers) GetFavoritesSchoolsHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(email)
	if err != nil {
		return err
	}

	schoolIDs, err := h.favoriteService.GetFavoriteSchoolIDs(user.ID)
	if err != nil {
		return err
	}

	schools, err := h.schoolService.GetSchoolsByIDs(schoolIDs)
	if err != nil {
		return err
	}

	return c.JSON(schools)
}

// @Summary Delete a favorite
// @Description Delete a favorite
// @Tags favorites
// @Accept json
// @Produce json
// @Param schoolID path int true "School ID"
// @Success 200 {object} string "OK"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /favorites/{schoolID} [delete]
func (h *FavoriteHandlers) DeleteFavoriteHandler(c *fiber.Ctx) error {
	token := c.Cookies("token")

	email, err := userDomain.GetEmailFromToken(token)
	if err != nil {
		return err
	}

	user, err := h.userService.GetByEmail(email)
	if err != nil {
		return err
	}

	schoolID, err := strconv.Atoi(c.Params("schoolID"))
	if err != nil {
		return err
	}

	err = h.favoriteService.DeleteFavorite(user.ID, schoolID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
