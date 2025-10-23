package routes

import (
	"ccpp-backend/internal/infrastructure/handlers"
	"ccpp-backend/internal/domain/services"
	infraServices "ccpp-backend/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userService services.UserService, facebookService *infraServices.FacebookService) {
	userHandler := handlers.NewUserHandler(userService)
	galleryHandler := handlers.NewGalleryHandler(facebookService)

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Gallery endpoints
		gallery := api.Group("/gallery")
		{
			gallery.GET("/photos", galleryHandler.GetPhotos)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "API is running",
		})
	})
}
