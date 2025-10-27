package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type contactRepositoryImpl struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) repositories.ContactRepository {
	return &contactRepositoryImpl{db: db}
}

func (r *contactRepositoryImpl) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

func (r *contactRepositoryImpl) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Order("created_at desc").Find(&contacts).Error
	return contacts, err
}

func (r *contactRepositoryImpl) FindByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.First(&contact, id).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *contactRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Contact{}, id).Error
}

