package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/config"
	"github.com/k61b/okul/internal/application/favoriteservice"
	"github.com/k61b/okul/internal/application/schoolservice"
	"github.com/k61b/okul/internal/application/userservice"
	"github.com/k61b/okul/internal/application/verificationservice"
	"github.com/k61b/okul/internal/infrastructure/database/postgres"
	"github.com/k61b/okul/internal/infrastructure/repository"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
	"github.com/k61b/okul/web/api/routes"

	_ "github.com/k61b/okul/docs"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Okul API
// @description This is a sample server Okul server.
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.LoadConfig("dev") // Adjust the environment ("dev", "prod", etc.) as needed
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Construct PostgreSQL database connection string
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	// Initialize PostgreSQL database
	db, err := postgres.NewPostgreSQLDB(dbConnectionString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewPostgresUserRepository(db.DB())
	schoolRepo := repository.NewPostgresSchoolRepository(db.DB())
	verificationRepo := repository.NewPostgresVerificationRepository(db.DB())
	favoriteRepo := repository.NewPostgresFavoriteRepository(db.DB())

	// Initialize application services
	userService := userservice.NewUserService(userRepo)
	schoolService := schoolservice.NewSchoolService(schoolRepo)
	verificationService := verificationservice.NewVerificationService(verificationRepo)
	favoriteService := favoriteservice.NewFavoriteService(favoriteRepo)

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(middleware.RequestLoggerMiddleware)
	app.Use(middleware.CORSMiddleware)

	// Initialize handlers
	schoolHandlers := handlers.NewSchoolHandlers(schoolService)
	userHandlers := handlers.NewUserHandlers(userService, verificationService)
	favoriteHandlers := handlers.NewFavoriteHandlers(favoriteService, schoolService, userService)

	// Initialize routes
	routes.SetupSchoolRoutes(app, schoolHandlers)
	routes.SetupUserRoutes(app, userHandlers)
	routes.SetupFavoriteRoutes(app, favoriteHandlers)

	// Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Start the Fiber app
	fmt.Printf("Server is running on port %s\n", cfg.Server.Port)
	app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
