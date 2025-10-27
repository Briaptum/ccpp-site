package routes

import (
	"ccpp-backend/internal/infrastructure/handlers"
	"ccpp-backend/internal/domain/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userService services.UserService, contactService services.ContactService) {
	userHandler := handlers.NewUserHandler(userService)
	contactHandler := handlers.NewContactHandler(contactService)

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

		contacts := api.Group("/contacts")
		{
			contacts.POST("", contactHandler.CreateContact)
			contacts.GET("", contactHandler.GetAllContacts)
			contacts.GET("/:id", contactHandler.GetContactByID)
			contacts.DELETE("/:id", contactHandler.DeleteContact)
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
