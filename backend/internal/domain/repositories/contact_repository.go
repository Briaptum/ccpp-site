package repositories

import "ccpp-backend/internal/domain/models"

type ContactRepository interface {
	Create(contact *models.Contact) error
	FindAll() ([]models.Contact, error)
	FindByID(id uint) (*models.Contact, error)
	Delete(id uint) error
}

