package entity

import "time"

type Ticket struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    EventID   uint      `json:"event_id"`
    UserID    uint      `json:"user_id"`
    Quantity  int       `json:"quantity"`
    TotalPrice float64  `json:"total_price"`
    Status    string    `json:"status"` // e.g., "active", "canceled"
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    Event Event `gorm:"foreignKey:EventID" json:"event"`
    User  User  `gorm:"foreignKey:UserID" json:"user"`
}
