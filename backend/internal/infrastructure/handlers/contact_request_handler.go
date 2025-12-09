package handlers

import (
	"log"
	"net/http"
	"strconv"

	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type ContactRequestHandler struct {
	service services.ContactRequestService
}

func NewContactRequestHandler(service services.ContactRequestService) *ContactRequestHandler {
	return &ContactRequestHandler{service: service}
}

// CreateContactRequest mirrors the Rolston HVAC contact endpoint
func (h *ContactRequestHandler) CreateContactRequest(c *gin.Context) {
	var req struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Reason  string `json:"reason" binding:"required"`
		Phone   string `json:"phone"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request. Name, email, and message are required.",
		})
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	request := models.ContactRequest{
		Name:      req.Name,
		Email:     req.Email,
		Reason:    req.Reason,
		Message:   req.Message,
		IPAddress: &ipAddress,
		UserAgent: &userAgent,
		Metadata:  models.JSONB{},
	}

	if req.Phone != "" {
		request.Phone = &req.Phone
	}

	if err := h.service.CreateContactRequest(&request); err != nil {
		log.Printf("Error creating contact request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save contact request",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contact request received successfully",
		"id":      request.ID,
	})
}

func (h *ContactRequestHandler) GetContactRequests(c *gin.Context) {
	requests, err := h.service.GetContactRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch contact requests",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requests": requests,
	})
}

func (h *ContactRequestHandler) GetContactRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact request ID"})
		return
	}

	request, err := h.service.GetContactRequest(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact request not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"request": request,
	})
}

func (h *ContactRequestHandler) DeleteContactRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact request ID"})
		return
	}

	if err := h.service.DeleteContactRequest(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact request deleted successfully"})
}
