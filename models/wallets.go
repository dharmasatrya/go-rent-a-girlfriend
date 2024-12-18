package models

import (
	"time"

	"gorm.io/gorm"
)

// Wallet represents wallet information
type Wallet struct {
	ID                uint           `json:"id" example:"1"`
	UserID            uint           `json:"user_id" example:"1" binding:"required"`
	BankCode          string         `json:"bank_code" example:"BCA" binding:"required"`
	BankAccountNumber string         `json:"bank_account_number" example:"1234567890" binding:"required"`
	BankAccountName   string         `json:"bank_account_name" example:"Dharma Satrya" binding:"required"`
	Balance           int            `json:"balance"`
	User              User           `json:"-" swaggerignore:"true"`
	CreatedAt         time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}
