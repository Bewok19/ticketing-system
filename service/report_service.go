package service

import (
	"errors"
	"ticketing-system/config"
	"ticketing-system/entity"
)

func GetSummaryReport() (*entity.SummaryReport, error) {
	var result entity.SummaryReport

	// Hitung total tiket terjual dan pendapatan
	query := `
		SELECT 
			COALESCE(SUM(t.quantity), 0) AS total_tickets,
			COALESCE(SUM(t.quantity * e.price), 0) AS total_revenue
		FROM tickets t
		JOIN events e ON t.event_id = e.id
		WHERE t.status = 'active'
	`

	if err := config.DB.Raw(query).Scan(&result).Error; err != nil {
		return nil, errors.New("failed to retrieve summary report")
	}

	return &result, nil
}

func GetEventReport(eventID uint) (*entity.EventReport, error) {
    var event entity.Event
    if err := config.DB.First(&event, eventID).Error; err != nil {
        return nil, errors.New("event not found")
    }

    var totalTickets int
    var totalRevenue float64

    if err := config.DB.Model(&entity.Ticket{}).
        Where("event_id = ?", eventID).
        Select("SUM(quantity) as total_tickets, SUM(total_price) as total_revenue").
        Row().Scan(&totalTickets, &totalRevenue); err != nil {
        return nil, errors.New("failed to generate report")
    }

    return &entity.EventReport{
        EventName:    event.Name,
        TotalTickets: totalTickets,
        TotalRevenue: totalRevenue,
    }, nil
}
