package services

import "ccpp-backend/internal/domain/models"

type ContactRequestService interface {
	CreateContactRequest(request *models.ContactRequest) error
	GetContactRequests() ([]models.ContactRequest, error)
	GetContactRequest(id uint) (*models.ContactRequest, error)
	DeleteContactRequest(id uint) error
}
