package entity

import "time"

type Event struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"unique;not null"`
    Capacity  int       `gorm:"not null"`
    Price     float64   `gorm:"not null"`
    Status    string    `gorm:"not null"`
    EventDate time.Time `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type EventInput struct {
    Name      string    `json:"name" binding:"required"`
    Capacity  int       `json:"capacity" binding:"required,gte=0"`
    Price     float64   `json:"price" binding:"required,gte=0"`
    Status    string    `json:"status" binding:"required"`
    EventDate time.Time `json:"event_date" binding:"required"`
}
