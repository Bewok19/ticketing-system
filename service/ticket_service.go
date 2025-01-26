package service

import (
	"errors"
	"ticketing-system/config"
	"ticketing-system/entity"
)

func CreateTicket(userID, eventID uint, quantity int) (*entity.Ticket, error) {
	var event entity.Event
	if err := config.DB.First(&event, eventID).Error; err != nil {
		return nil, errors.New("event not found")
	}

	if event.Capacity < quantity {
		return nil, errors.New("not enough capacity")
	}

	event.Capacity -= quantity
	if err := config.DB.Save(&event).Error; err != nil {
		return nil, errors.New("failed to update event capacity")
	}

	ticket := entity.Ticket{
		EventID:    eventID,
		UserID:     userID,
		Quantity:   quantity,
		TotalPrice: float64(quantity) * event.Price,
		Status:     "active",
	}
	if err := config.DB.Create(&ticket).Error; err != nil {
		return nil, errors.New("failed to create ticket")
	}

	return &ticket, nil
}

// GetTickets retrieves tickets for a user (or all if admin)
func GetTickets(userID uint, userRole string) ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	query := config.DB.Debug().Preload("Event")

	// If the user is not an admin, filter by userID
	if userRole != "admin" {
		query = query.Where("user_id = ?", userID)
	}

	// Fetch tickets from the database
	if err := query.Find(&tickets).Error; err != nil {
		return nil, errors.New("failed to fetch tickets")
	}

	return tickets, nil
}



func DeleteTicket(ticketID, userID uint, isAdmin bool) error {
	var ticket entity.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		return errors.New("ticket not found")
	}

	if !isAdmin && ticket.UserID != userID {
		return errors.New("unauthorized")
	}

	ticket.Status = "canceled"
	if err := config.DB.Save(&ticket).Error; err != nil {
		return errors.New("failed to cancel ticket")
	}

	return nil
}
