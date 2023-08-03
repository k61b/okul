package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/k61b/okul/config"
	"github.com/k61b/okul/internal/application/adminservice"
	"github.com/k61b/okul/internal/application/schoolservice"
	"github.com/k61b/okul/internal/application/userservice"
	"github.com/k61b/okul/internal/infrastructure/database/postgres"
	"github.com/k61b/okul/web/api/handlers"
	"github.com/k61b/okul/web/api/middleware"
	"github.com/k61b/okul/web/api/routes"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize PostgreSQL database
	db, err := postgres.NewPostgreSQLDB(cfg.GetString("database.connection_string"))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := postgres.NewPostgresUserRepository(db.DB())
	schoolRepo := postgres.NewPostgresSchoolRepository(db.DB())
	adminRepo := postgres.NewPostgresAdminRepository(db.DB()) // Implement this

	// Initialize application services
	userService := userservice.NewUserService(userRepo)
	schoolService := schoolservice.NewSchoolService(schoolRepo)
	adminService := adminservice.NewAdminService(adminRepo) // Implement this

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(middleware.RequestLoggerMiddleware)
	app.Use(middleware.CORSMiddleware)

	// Initialize handlers
	adminHandlers := handlers.NewAdminHandlers(adminService)
	schoolHandlers := handlers.NewSchoolHandlers(schoolService)
	userHandlers := handlers.NewUserHandlers(userService)

	// Initialize routes
	routes.SetupAdminRoutes(app, adminHandlers)
	routes.SetupSchoolRoutes(app, schoolHandlers)
	routes.SetupUserRoutes(app, userHandlers)

	// Start the Fiber app
	port := cfg.GetInt("server.port")
	fmt.Printf("Server is running on port %d\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
