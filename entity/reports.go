package entity

type SummaryReport struct {
	TotalTickets int     `json:"total_tickets"`
	TotalRevenue float64 `json:"total_revenue"`
}

type EventReport struct {
    EventName    string  `json:"event_name"`
    TotalTickets int     `json:"total_tickets"`
    TotalRevenue float64 `json:"total_revenue"`
}
