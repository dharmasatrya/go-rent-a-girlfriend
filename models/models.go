package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
// @Description User account information
// swagger:model User
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"not null" json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	Email     string         `gorm:"not null" json:"email"`
	Role      string         `gorm:"not null" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Wallet represent the user's wallet
// @Description Wallet information
// swagger:model Wallet
type Wallet struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	UserID            uint           `gorm:"not null" json:"user_id"`
	BankCode          string         `gorm:"not null" json:"bank_code"`
	BankAccountNumber string         `gorm:"not null" json:"bank_account_number"`
	User              User           `gorm:"foreignKey:UserID"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Boy represents the boy profile model
// @Description Boy profile information
// swagger:model Boy
type Boy struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	UserID            uint           `json:"user_id" gorm:"not null"`
	FirstName         string         `gorm:"not null" json:"first_name"`
	LastName          string         `gorm:"not null" json:"last_name"`
	Age               int            `gorm:"not null" json:"age"`
	ProfilePictureURL string         `json:"profile_picture_url"`
	Bio               string         `json:"bio"`
	User              User           `gorm:"foreignKey:UserID"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Girl represents the girl profile model
// @Description girl profile information
// swagger:model Girl
type Girl struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	UserID            uint           `gorm:"not null" json:"user_id"`
	FirstName         string         `gorm:"not null" json:"first_name"`
	LastName          string         `gorm:"not null" json:"last_name"`
	Age               int            `gorm:"not null" json:"age"`
	ProfilePictureURL string         `json:"profile_picture_url"`
	Bio               string         `json:"bio"`
	DailyRate         int            `gorm:"not null" json:"daily_rate"`
	User              User           `gorm:"foreignKey:UserID"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Availability represents the availability of the girls
// @Description girls availability
// swagger:model Availability
type Availability struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	GirlID      uint           `gorm:"not null" json:"girl_id"`
	IsAvailable bool           `gorm:"not null" json:"is_available"`
	Girl        Girl           `gorm:"foreignKey:GirlID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Transaction
// @Description transactions table model
// swagger:model Transaction
type Transaction struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	SenderWalletID   uint           `gorm:"not null" json:"sender_wallet_id"`
	ReceiverWalletID uint           `gorm:"not null" json:"receiver_wallet_id"`
	Amount           int            `gorm:"not null" json:"amount"`
	TransactionDate  time.Time      `gorm:"not null" json:"transaction_date"`
	SenderWallet     Wallet         `gorm:"foreignKey:SenderWalletID"`
	ReceiverWallet   Wallet         `gorm:"foreignKey:ReceiverWalletID"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Booking
// @Description booking table model
// swagger:model Booking
type Booking struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	BoyID       uint           `gorm:"not null" json:"boy_id"`
	GirlID      uint           `gorm:"not null" json:"girl_id"`
	BookingDate time.Time      `gorm:"not null" json:"booking_date"`
	NumOfDays   int            `gorm:"not null" json:"num_of_days"`
	TotalCost   int            `gorm:"not null" json:"total_cost"`
	Boy         Boy            `gorm:"foreignKey:BoyID"`
	Girl        Girl           `gorm:"foreignKey:GirlID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Rating
// @Description rating
// swagger:model Rating
type Rating struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	GirlID    uint           `gorm:"not null" json:"girl_id"`
	Review    string         `json:"review"`
	Stars     int            `gorm:"not null;check:stars >= 1 AND stars <= 5" json:"stars"`
	Girl      Girl           `gorm:"foreignKey:GirlID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
