package main

import (
	"log"
	"os"

	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/infrastructure/database"
	infraRepos "ccpp-backend/internal/infrastructure/repositories"
	infraServices "ccpp-backend/internal/infrastructure/services"
	"ccpp-backend/internal/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	database.InitDB()
	db := database.GetDB()

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.User{}, &models.Contact{}, &models.Gallery{}, &models.Event{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := infraRepos.NewUserRepository(db)
	contactRepo := infraRepos.NewContactRepository(db)
	galleryRepo := infraRepos.NewGalleryRepository(db)
	eventRepo := infraRepos.NewEventRepository(db)

	// Initialize services
	userService := infraServices.NewUserService(userRepo)
	contactService := infraServices.NewContactService(contactRepo)
	galleryService := infraServices.NewGalleryService(galleryRepo)
	eventService := infraServices.NewEventService(eventRepo)

	// Setup Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Setup routes
	routes.SetupRoutes(router, userService, contactService, galleryService, eventService)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
