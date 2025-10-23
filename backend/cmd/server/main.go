package main

import (
	"log"
	"os"

	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"
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
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := infraRepos.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	
	// Initialize Facebook service
	facebookAccessToken := os.Getenv("FACEBOOK_ACCESS_TOKEN")
	if facebookAccessToken == "" {
		log.Println("Warning: FACEBOOK_ACCESS_TOKEN not set, gallery features will not work")
	}
	facebookService := infraServices.NewFacebookService(facebookAccessToken)

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
	routes.SetupRoutes(router, userService, facebookService)

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
