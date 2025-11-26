package repositories

import "ccpp-backend/internal/domain/models"

type GalleryRepository interface {
	Create(gallery *models.Gallery) error
	FindAll() ([]models.Gallery, error)
	FindByCategory(category string) ([]models.Gallery, error)
	FindByID(id uint) (*models.Gallery, error)
	Delete(id uint) error
}

