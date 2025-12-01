package repositories

import "ccpp-backend/internal/domain/models"

type EventRepository interface {
	Create(event *models.Event) error
	FindAll() ([]models.Event, error)
	FindUpcoming() ([]models.Event, error)
	FindByID(id uint) (*models.Event, error)
	Update(event *models.Event) error
	Delete(id uint) error
}


