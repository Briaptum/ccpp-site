package handlers

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GalleryHandler struct {
	galleryService services.GalleryService
	uploadPath     string
}

func NewGalleryHandler(galleryService services.GalleryService) *GalleryHandler {
	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "./uploads/gallery"
	}

	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		log.Printf("Warning: Failed to create upload directory: %v", err)
	}

	return &GalleryHandler{
		galleryService: galleryService,
		uploadPath:     uploadPath,
	}
}

// UploadGallery handles POST /api/gallery/upload
func (h *GalleryHandler) UploadGallery(c *gin.Context) {
	// Get category from form
	category := c.PostForm("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	// Get the file from form
	file, err := c.FormFile("image")
	if err != nil {
		log.Printf("Error getting file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif"}
	isValid := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			isValid = true
			break
		}
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only JPG, PNG, and GIF are allowed"})
		return
	}

	// Validate file size (10MB max)
	if file.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB limit"})
		return
	}

	// Generate unique filename
	timestamp := time.Now().Unix()
	filename := filepath.Base(file.Filename)
	nameWithoutExt := strings.TrimSuffix(filename, ext)
	uniqueFilename := fmt.Sprintf("%s_%d%s", nameWithoutExt, timestamp, ext)

	// Create category subdirectory
	categoryPath := filepath.Join(h.uploadPath, category)
	if err := os.MkdirAll(categoryPath, 0755); err != nil {
		log.Printf("Error creating category directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Save file
	filePath := filepath.Join(categoryPath, uniqueFilename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Printf("Error saving file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create gallery record
	gallery := &models.Gallery{
		Filename: uniqueFilename,
		Path:     filepath.Join("uploads/gallery", category, uniqueFilename),
		Category: category,
	}

	if err := h.galleryService.CreateGallery(gallery); err != nil {
		// If database save fails, delete the uploaded file
		os.Remove(filePath)
		log.Printf("Error creating gallery record: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save gallery record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Image uploaded successfully",
		"gallery": gallery,
	})
}

// GetAllGalleries handles GET /api/gallery
func (h *GalleryHandler) GetAllGalleries(c *gin.Context) {
	category := c.Query("category")
	var galleries []models.Gallery
	var err error

	if category != "" {
		galleries, err = h.galleryService.GetGalleriesByCategory(category)
	} else {
		galleries, err = h.galleryService.GetAllGalleries()
	}

	if err != nil {
		log.Printf("Error getting galleries: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve galleries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"galleries": galleries,
		"count":     len(galleries),
	})
}

// GetGalleryByID handles GET /api/gallery/:id
func (h *GalleryHandler) GetGalleryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gallery ID"})
		return
	}

	gallery, err := h.galleryService.GetGalleryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}

	c.JSON(http.StatusOK, gallery)
}

// DeleteGallery handles DELETE /api/gallery/:id
func (h *GalleryHandler) DeleteGallery(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gallery ID"})
		return
	}

	// Get gallery record first to get file path
	gallery, err := h.galleryService.GetGalleryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gallery not found"})
		return
	}

	// Delete file from filesystem
	filePath := filepath.Join(h.uploadPath, gallery.Category, gallery.Filename)
	if err := os.Remove(filePath); err != nil {
		log.Printf("Warning: Failed to delete file %s: %v", filePath, err)
		// Continue with database deletion even if file deletion fails
	}

	// Delete from database
	if err := h.galleryService.DeleteGallery(uint(id)); err != nil {
		log.Printf("Error deleting gallery: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete gallery"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery deleted successfully"})
}

