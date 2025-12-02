package services

import "ccpp-backend/internal/domain/models"

type EventService interface {
	CreateEvent(event *models.Event) error
	GetAllEvents() ([]models.Event, error)
	GetUpcomingEvents() ([]models.Event, error)
	GetEventByID(id uint) (*models.Event, error)
	UpdateEvent(event *models.Event) error
	DeleteEvent(id uint) error
}




