package transaction

import (
	"errors"
	"ticketing-system/config"
	"ticketing-system/entity"

	"gorm.io/gorm"
)

// ProcessTransaction handles ticket creation and updates event capacity atomically
func ProcessTransaction(userID, eventID uint, quantity int) (*entity.Ticket, error) {
	// Start a new database transaction
	tx := config.DB.Begin()

	// Find event
	var event entity.Event
	if err := tx.First(&event, eventID).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("event not found")
	}

	// Check if capacity is sufficient
	if event.Capacity < quantity {
		tx.Rollback()
		return nil, errors.New("not enough capacity")
	}

	// Deduct capacity
	if err := tx.Model(&entity.Event{}).Where("id = ? AND capacity >= ?", eventID, quantity).
		UpdateColumn("capacity", gorm.Expr("capacity - ?", quantity)).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to update event capacity")
	}

	// Create ticket
	ticket := entity.Ticket{
		EventID:    eventID,
		UserID:     userID,
		Quantity:   quantity,
		TotalPrice: float64(quantity) * event.Price,
		Status:     "active",
	}
	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create ticket")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &ticket, nil
}
