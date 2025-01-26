package repository

import (
	"ticketing-system/entity"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

// FindByName retrieves an event by its name
func (r *EventRepository) FindByName(name string) (*entity.Event, error) {
	var event entity.Event
	err := r.db.Where("LOWER(name) = ?", name).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Save saves a new event to the database
func (r *EventRepository) Save(event *entity.Event) error {
	return r.db.Create(event).Error
}
