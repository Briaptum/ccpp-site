package routes

import (
	"ccpp-backend/internal/infrastructure/handlers"
	"ccpp-backend/internal/domain/services"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userService services.UserService, contactService services.ContactService, galleryService services.GalleryService, eventService services.EventService) {
	userHandler := handlers.NewUserHandler(userService)
	contactHandler := handlers.NewContactHandler(contactService)
	galleryHandler := handlers.NewGalleryHandler(galleryService)
	eventHandler := handlers.NewEventHandler(eventService)

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

		gallery := api.Group("/gallery")
		{
			gallery.POST("/upload", galleryHandler.UploadGallery)
			gallery.GET("", galleryHandler.GetAllGalleries)
			gallery.GET("/:id", galleryHandler.GetGalleryByID)
			gallery.DELETE("/:id", galleryHandler.DeleteGallery)
		}

		events := api.Group("/events")
		{
			events.POST("", eventHandler.CreateEvent)
			events.GET("", eventHandler.GetAllEvents)
			events.GET("/:id", eventHandler.GetEventByID)
			events.PUT("/:id", eventHandler.UpdateEvent)
			events.DELETE("/:id", eventHandler.DeleteEvent)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "API is running",
		})
	})

	// Serve static files for uploads
	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "./uploads"
	}
	router.Static("/uploads", uploadPath)
}
