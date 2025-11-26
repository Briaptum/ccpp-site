package repositories

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type galleryRepositoryImpl struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) repositories.GalleryRepository {
	return &galleryRepositoryImpl{db: db}
}

func (r *galleryRepositoryImpl) Create(gallery *models.Gallery) error {
	return r.db.Create(gallery).Error
}

func (r *galleryRepositoryImpl) FindAll() ([]models.Gallery, error) {
	var galleries []models.Gallery
	err := r.db.Order("created_at desc").Find(&galleries).Error
	return galleries, err
}

func (r *galleryRepositoryImpl) FindByCategory(category string) ([]models.Gallery, error) {
	var galleries []models.Gallery
	err := r.db.Where("category = ?", category).Order("created_at desc").Find(&galleries).Error
	return galleries, err
}

func (r *galleryRepositoryImpl) FindByID(id uint) (*models.Gallery, error) {
	var gallery models.Gallery
	err := r.db.First(&gallery, id).Error
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (r *galleryRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Gallery{}, id).Error
}

