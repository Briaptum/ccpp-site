package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type contactRequestRepositoryImpl struct {
	db *gorm.DB
}

func NewContactRequestRepository(db *gorm.DB) repositories.ContactRequestRepository {
	return &contactRequestRepositoryImpl{db: db}
}

func (r *contactRequestRepositoryImpl) Create(request *models.ContactRequest) error {
	return r.db.Create(request).Error
}

func (r *contactRequestRepositoryImpl) FindAll() ([]models.ContactRequest, error) {
	var requests []models.ContactRequest
	err := r.db.Order("created_at desc").Find(&requests).Error
	return requests, err
}

func (r *contactRequestRepositoryImpl) FindByID(id uint) (*models.ContactRequest, error) {
	var request models.ContactRequest
	err := r.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *contactRequestRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.ContactRequest{}, id).Error
}
