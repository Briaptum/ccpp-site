package services

import (
	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"ccpp-backend/internal/domain/services"
)

type eventServiceImpl struct {
	eventRepo repositories.EventRepository
}

func NewEventService(eventRepo repositories.EventRepository) services.EventService {
	return &eventServiceImpl{eventRepo: eventRepo}
}

func (s *eventServiceImpl) CreateEvent(event *models.Event) error {
	return s.eventRepo.Create(event)
}

func (s *eventServiceImpl) GetAllEvents() ([]models.Event, error) {
	return s.eventRepo.FindAll()
}

func (s *eventServiceImpl) GetUpcomingEvents() ([]models.Event, error) {
	return s.eventRepo.FindUpcoming()
}

func (s *eventServiceImpl) GetEventByID(id uint) (*models.Event, error) {
	return s.eventRepo.FindByID(id)
}

func (s *eventServiceImpl) UpdateEvent(event *models.Event) error {
	return s.eventRepo.Update(event)
}

func (s *eventServiceImpl) DeleteEvent(id uint) error {
	return s.eventRepo.Delete(id)
}




