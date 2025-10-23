package handlers

import (
	"net/http"
	"strconv"

	"ccpp-backend/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
)

type GalleryHandler struct {
	facebookService *services.FacebookService
}

func NewGalleryHandler(facebookService *services.FacebookService) *GalleryHandler {
	return &GalleryHandler{
		facebookService: facebookService,
	}
}

func (h *GalleryHandler) GetPhotos(c *gin.Context) {
	// Get limit from query params, default to 12
	limitStr := c.DefaultQuery("limit", "12")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 12
	}

	// Facebook Page ID for "Suki Lv"
	pageID := "111245058546149"

	// Fetch photos from Facebook page
	photos, err := h.facebookService.GetPagePhotos(pageID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch photos from Facebook",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": photos,
		"count":  len(photos),
	})
}
