package handlers

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	contactService services.ContactService
}

func NewContactHandler(contactService services.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

// CreateContact handles POST /api/contacts
func (h *ContactHandler) CreateContact(c *gin.Context) {
	var contact models.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save contact"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contact submitted successfully",
		"contact": contact,
	})
}

// GetAllContacts handles GET /api/contacts
func (h *ContactHandler) GetAllContacts(c *gin.Context) {
	contacts, err := h.contactService.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve contacts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
		"count":    len(contacts),
	})
}

// GetContactByID handles GET /api/contacts/:id
func (h *ContactHandler) GetContactByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	contact, err := h.contactService.GetContactByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// DeleteContact handles DELETE /api/contacts/:id
func (h *ContactHandler) DeleteContact(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contact ID"})
		return
	}

	if err := h.contactService.DeleteContact(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}

