package response

import (
	"ticketing-system/entity"
	"time"
)

// Struct untuk format respons umum
type GeneralResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"` // Field opsional untuk data
}

// Struct spesifik untuk Event
type EventResponse struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Capacity  int       `json:"capacity"`
    Price     float64   `json:"price"`
    Status    string    `json:"status"`
    EventDate time.Time `json:"event_date"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Struct untuk respons daftar event
type EventListResponse struct {
    ID       uint    `json:"id"`
    Name     string  `json:"name"`
    Capacity int     `json:"capacity"`
    Price    float64 `json:"price"`
    Status   string  `json:"status"`
}

type TicketResponse struct {
	ID         uint   `json:"id"`
	EventName  string `json:"event_name"`
	UserID     uint   `json:"user_id"`
	Quantity   int    `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string `json:"status"`
}

// ToTicketResponses converts a slice of Ticket entities to TicketResponses
func ToTicketResponses(tickets []entity.Ticket) []TicketResponse {
	var responses []TicketResponse
	for _, ticket := range tickets {
		responses = append(responses, TicketResponse{
			ID:         ticket.ID,
			EventName:  ticket.Event.Name,
			UserID:     ticket.UserID,
			Quantity:   ticket.Quantity,
			TotalPrice: ticket.TotalPrice,
			Status:     ticket.Status,
		})
	}
	return responses
}

