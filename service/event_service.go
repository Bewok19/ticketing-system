package service

import (
	"errors"
	"strings"
	"ticketing-system/config"
	"ticketing-system/entity"

	"ticketing-system/repository"
)

type EventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

// CheckEventName checks if the event name already exists in the database
func (s *EventService) CheckEventName(name string) error {
	existingEvent, err := s.repo.FindByName(strings.ToLower(name))
	if err == nil && existingEvent.ID != 0 {
		return errors.New("event name already exists")
	}
	return nil
}

func (s *EventService) CreateEvent(event entity.Event) (*entity.Event, error) {
	// Periksa apakah nama event sudah ada di database
	var existingEvent entity.Event
	if err := config.DB.Where("name = ?", event.Name).First(&existingEvent).Error; err == nil {
		// Jika nama event sudah ada, kembalikan error
		return nil, errors.New("event name already exists")
	}

	// Jika nama belum ada, lanjutkan dengan menyimpan data
	if err := config.DB.Create(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

