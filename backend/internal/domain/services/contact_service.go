package services

import "ccpp-backend/internal/domain/models"

type ContactService interface {
	CreateContact(contact *models.Contact) error
	GetAllContacts() ([]models.Contact, error)
	GetContactByID(id uint) (*models.Contact, error)
	DeleteContact(id uint) error
}

