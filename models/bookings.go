package models

import (
	"time"

	"gorm.io/gorm"
)

// Availability represents availability information
type Availability struct {
	ID          uint           `json:"id" example:"1"`
	GirlID      uint           `json:"girl_id" example:"1" binding:"required"`
	IsAvailable bool           `json:"is_available" example:"true" binding:"required"`
	Girl        Girl           `json:"-" swaggerignore:"true"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	CreatedAt   time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

// Booking represents booking information
type Booking struct {
	ID          uint           `json:"id" example:"1"`
	BoyID       uint           `json:"boy_id" example:"1" binding:"required"`
	GirlID      uint           `json:"girl_id" example:"1" binding:"required"`
	BookingDate time.Time      `json:"booking_date" example:"2024-01-01T00:00:00Z" binding:"required"`
	NumOfDays   int            `json:"num_of_days" example:"3" binding:"required"`
	TotalCost   int            `json:"total_cost" example:"300" binding:"required"`
	Boy         Boy            `json:"boy,omitempty"`
	Girl        Girl           `json:"girl,omitempty"`
	CreatedAt   time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

type BookingRequest struct {
	GirlID      uint      `json:"girl_id" example:"1" binding:"required"`
	BookingDate time.Time `json:"booking_date" example:"2024-01-01T00:00:00Z" binding:"required"`
	NumOfDays   int       `json:"num_of_days" example:"3" binding:"required"`
}
