package models

import (
	"time"

	"gorm.io/gorm"
)

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
