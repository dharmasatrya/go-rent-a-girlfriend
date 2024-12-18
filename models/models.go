package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents user information
type User struct {
	ID        uint           `json:"id" example:"1"`
	Username  string         `json:"username" example:"johndoe" binding:"required"`
	Password  string         `json:"-" example:"password123" binding:"required"`
	Email     string         `json:"email" example:"john@example.com" binding:"required"`
	Role      string         `json:"role" example:"boy" binding:"required"`
	CreatedAt time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

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

// Boy represents boy profile information
type Boy struct {
	ID                uint           `json:"id" example:"1"`
	UserID            uint           `json:"user_id" example:"1" binding:"required"`
	FirstName         string         `json:"first_name" example:"John" binding:"required"`
	LastName          string         `json:"last_name" example:"Doe" binding:"required"`
	Age               int            `json:"age" example:"25" binding:"required"`
	ProfilePictureURL string         `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	Bio               string         `json:"bio" example:"I love traveling and meeting new people"`
	User              User           `json:"-" swaggerignore:"true"`
	CreatedAt         time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

// Girl represents girl profile information
type Girl struct {
	ID                uint           `json:"id" example:"1"`
	UserID            uint           `json:"user_id" example:"1" binding:"required"`
	FirstName         string         `json:"first_name" example:"Jane" binding:"required"`
	LastName          string         `json:"last_name" example:"Doe" binding:"required"`
	Age               int            `json:"age" example:"23" binding:"required"`
	ProfilePictureURL string         `json:"profile_picture_url" example:"https://example.com/profile.jpg"`
	Bio               string         `json:"bio" example:"I enjoy meeting new people"`
	DailyRate         int            `json:"daily_rate" example:"100" binding:"required"`
	User              User           `json:"-" swaggerignore:"true"`
	CreatedAt         time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

// Availability represents availability information
type Availability struct {
	ID          uint           `json:"id" example:"1"`
	GirlID      uint           `json:"girl_id" example:"1" binding:"required"`
	IsAvailable bool           `json:"is_available" example:"true" binding:"required"`
	Girl        Girl           `json:"-" swaggerignore:"true"`
	CreatedAt   time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

// Transaction represents transaction information
type Transaction struct {
	ID               uint           `json:"id" example:"1"`
	SenderWalletID   uint           `json:"sender_wallet_id" example:"1" binding:"required"`
	ReceiverWalletID uint           `json:"receiver_wallet_id" example:"2" binding:"required"`
	Amount           int            `json:"amount" example:"1000" binding:"required"`
	TransactionDate  time.Time      `json:"transaction_date" example:"2024-01-01T00:00:00Z" binding:"required"`
	SenderWallet     Wallet         `json:"-" swaggerignore:"true"`
	ReceiverWallet   Wallet         `json:"-" swaggerignore:"true"`
	CreatedAt        time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt        time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
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

// Rating represents rating information
type Rating struct {
	ID        uint           `json:"id" example:"1"`
	GirlID    uint           `json:"girl_id" example:"1" binding:"required"`
	Review    string         `json:"review" example:"Great experience!"`
	Stars     int            `json:"stars" example:"5" binding:"required"`
	Girl      Girl           `json:"-" swaggerignore:"true"`
	CreatedAt time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

// LoginRequest represents login request information
type LoginRequest struct {
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"password123"`
}

type BookingRequest struct {
	GirlID      uint      `json:"girl_id" example:"1" binding:"required"`
	BookingDate time.Time `json:"booking_date" example:"2024-01-01T00:00:00Z" binding:"required"`
	NumOfDays   int       `json:"num_of_days" example:"3" binding:"required"`
}

type XenditInvoiceRequest struct {
	ExternalId  string
	Amount      int
	Description string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
}

type DepostitRequest struct {
	Amount int
}

type InternalTransaction struct {
	ID         uint           `json:"id"`
	UserID     uint           `json:"user_id"`
	ExternalId string         `json:"external_id"`
	Amount     int            `json:"amount"`
	Status     string         `json:"status"`
	Type       string         `json:"type"`
	CreatedAt  time.Time      `json:"created_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty" swaggerignore:"true" example:"2024-01-01T00:00:00Z"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" swaggerignore:"true" swaggertype:"string" example:"2024-01-01T00:00:00Z"`
}

type XenditCallback struct {
	ExternalId string `json:"external_id"`
	Status     string `json:"status"`
}

type WithdrawalRequest struct {
	Amount int
}

type XenditDisbursementRequest struct {
	ExternalId        string
	Amount            int
	BankCode          string
	AccountHolderName string
	BankAccountNumber string
	Description       string
	Email             string
}
