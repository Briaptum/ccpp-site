package services

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"ccpp-backend/internal/domain/services"
)

type contactServiceImpl struct {
	contactRepo repositories.ContactRepository
}

func NewContactService(contactRepo repositories.ContactRepository) services.ContactService {
	return &contactServiceImpl{contactRepo: contactRepo}
}

func (s *contactServiceImpl) CreateContact(contact *models.Contact) error {
	return s.contactRepo.Create(contact)
}

func (s *contactServiceImpl) GetAllContacts() ([]models.Contact, error) {
	return s.contactRepo.FindAll()
}

func (s *contactServiceImpl) GetContactByID(id uint) (*models.Contact, error) {
	return s.contactRepo.FindByID(id)
}

func (s *contactServiceImpl) DeleteContact(id uint) error {
	return s.contactRepo.Delete(id)
}

