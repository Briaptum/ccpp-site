package repositories

import "ccpp-backend/internal/domain/models"

type ContactRequestRepository interface {
	Create(request *models.ContactRequest) error
	FindAll() ([]models.ContactRequest, error)
	FindByID(id uint) (*models.ContactRequest, error)
	Delete(id uint) error
}
