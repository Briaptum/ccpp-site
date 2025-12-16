package routes

import (
	"os"

	"ccpp-backend/internal/domain/services"
	"ccpp-backend/internal/infrastructure/handlers"
	"ccpp-backend/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userService services.UserService,
	contactService services.ContactService,
	contactRequestService services.ContactRequestService,
	galleryService services.GalleryService,
	eventService services.EventService,
) {
	userHandler := handlers.NewUserHandler(userService)
	contactHandler := handlers.NewContactHandler(contactService)
	contactRequestHandler := handlers.NewContactRequestHandler(contactRequestService)
	galleryHandler := handlers.NewGalleryHandler(galleryService)
	eventHandler := handlers.NewEventHandler(eventService)
	authHandler := handlers.NewAuthHandler()

	// Legacy-free routes copied from Rolston HVAC for contact requests
	contactRequests := router.Group("/api/contact-requests")
	{
		contactRequests.POST("", contactRequestHandler.CreateContactRequest)
	}
	legacyProtected := contactRequests.Group("")
	legacyProtected.Use(middleware.Auth0Middleware())
	{
		legacyProtected.GET("", contactRequestHandler.GetContactRequests)
		legacyProtected.GET("/:id", contactRequestHandler.GetContactRequest)
		legacyProtected.DELETE("/:id", contactRequestHandler.DeleteContactRequest)
	}

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.GET("/logout", authHandler.Logout)
			auth.GET("/profile", middleware.Auth0Middleware(), authHandler.Profile)
		}

		protected := api.Group("")
		protected.Use(middleware.Auth0Middleware())

		users := protected.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		contactRequestsV1 := api.Group("/contact-requests")
		{
			contactRequestsV1.POST("", contactRequestHandler.CreateContactRequest)
			protectedContactRequests := contactRequestsV1.Group("")
			protectedContactRequests.Use(middleware.Auth0Middleware())
			{
				protectedContactRequests.GET("", contactRequestHandler.GetContactRequests)
				protectedContactRequests.GET("/:id", contactRequestHandler.GetContactRequest)
				protectedContactRequests.DELETE("/:id", contactRequestHandler.DeleteContactRequest)
			}
		}

		contacts := protected.Group("/contacts")
		{
			contacts.POST("", contactHandler.CreateContact)
			contacts.GET("", contactHandler.GetAllContacts)
			contacts.GET("/:id", contactHandler.GetContactByID)
			contacts.DELETE("/:id", contactHandler.DeleteContact)
		}

		gallery := api.Group("/gallery")
		{
			gallery.POST("/upload", middleware.Auth0Middleware(), galleryHandler.UploadGallery)
			gallery.GET("", galleryHandler.GetAllGalleries)
			gallery.GET("/:id", galleryHandler.GetGalleryByID)
			gallery.DELETE("/:id", middleware.Auth0Middleware(), galleryHandler.DeleteGallery)
		}

		events := api.Group("/events")
		{
			events.POST("", middleware.Auth0Middleware(), eventHandler.CreateEvent)
			events.GET("", eventHandler.GetAllEvents)
			events.GET("/:id", eventHandler.GetEventByID)
			events.PUT("/:id", middleware.Auth0Middleware(), eventHandler.UpdateEvent)
			events.DELETE("/:id", middleware.Auth0Middleware(), eventHandler.DeleteEvent)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
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
