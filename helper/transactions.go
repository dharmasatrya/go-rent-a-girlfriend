package helper

import (
	"fmt"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/models"
	"time"
)

func GetWalletIDByUserID(userID uint) (uint, error) {
	var wallet models.Wallet

	if err := db.GormDB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return 0, fmt.Errorf("wallet not found for user ID %d: %v", userID, err)
	}

	return wallet.ID, nil
}

func Transaction(sender_id, recipient_id uint, amount int) error {
	// Start a database transaction
	tx := db.GormDB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	// Get sender's wallet
	var senderWallet models.Wallet
	if err := tx.Where("id = ?", sender_id).First(&senderWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch sender wallet: %v", err)
	}

	// Get recipient's wallet
	var recipientWallet models.Wallet
	if err := tx.Where("id = ?", recipient_id).First(&recipientWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch recipient wallet: %v", err)
	}

	// Check if sender has sufficient balance
	if senderWallet.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance: required %d, available %d", amount, senderWallet.Balance)
	}

	// Deduct from sender
	if err := tx.Model(&senderWallet).Update("balance", senderWallet.Balance-amount).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to deduct from sender: %v", err)
	}

	// Add to recipient
	if err := tx.Model(&recipientWallet).Update("balance", recipientWallet.Balance+amount).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to add to recipient: %v", err)
	}

	// Create transaction record
	transaction := models.Transaction{
		SenderWalletID:   sender_id,
		ReceiverWalletID: recipient_id,
		Amount:           amount,
		TransactionDate:  time.Now(),
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create transaction record: %v", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
