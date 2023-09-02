package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
)

func SetupFavoriteRoutes(app *fiber.App, favoriteHandlers *handlers.FavoriteHandlers) {
	favorite := app.Group("/favorite")

	favorite.Post("/:schoolID", middleware.AuthMiddleware, favoriteHandlers.CreateFavoriteHandler)
	favorite.Get("/", middleware.AuthMiddleware, favoriteHandlers.GetFavoritesSchoolsHandler)
	favorite.Delete("/:schoolID", middleware.AuthMiddleware, favoriteHandlers.DeleteFavoriteHandler)
}
