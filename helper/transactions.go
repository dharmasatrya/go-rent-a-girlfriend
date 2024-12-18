package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func CreateXenditInvoice(invoiceReq models.XenditInvoiceRequest) (map[string]interface{}, error) {
	// Prepare the request payload
	xenditUrl := os.Getenv("XENDIT_INVOICE_URL")
	payload := map[string]interface{}{
		"external_id": invoiceReq.ExternalId,
		"amount":      invoiceReq.Amount,
		"description": invoiceReq.Description,
		"customer": map[string]interface{}{
			"given_names":   invoiceReq.FirstName,
			"surname":       invoiceReq.LastName,
			"email":         invoiceReq.Email,
			"mobile_number": invoiceReq.Phone,
		},
		"customer_notification_preference": map[string]interface{}{
			"invoice_created": []string{"whatsapp", "email"},
			"invoice_paid":    []string{"whatsapp", "email"},
		},
		"currency": "IDR",
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Create the request
	request, err := http.NewRequest("POST", xenditUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Get API key from environment variable and encode it
	apiKey := os.Getenv("XENDIT_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("XENDIT_API_KEY not found in environment variables")
	}
	encodedKey := base64.StdEncoding.EncodeToString([]byte(apiKey))

	// Set headers
	request.Header.Set("Authorization", "Basic "+encodedKey)
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	// Check if response indicates an error
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API request failed with status %d: %v", resp.StatusCode, result)
	}

	return result, nil
}

func CreateXenditDisbursement(disbursementReq models.XenditDisbursementRequest) (map[string]interface{}, error) {
	// Prepare the request payload
	xenditUrl := os.Getenv("XENDIT_DISBURSEMENT_URL")
	payload := map[string]interface{}{
		"external_id":         disbursementReq.ExternalId,
		"amount":              disbursementReq.Amount,
		"bank_code":           disbursementReq.BankCode,
		"account_holder_name": disbursementReq.AccountHolderName,
		"account_number":      disbursementReq.BankAccountNumber,
		"description":         disbursementReq.Description,
		"email_to":            []string{disbursementReq.Email},
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Create the request
	request, err := http.NewRequest("POST", xenditUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Get API key from environment variable and encode it
	apiKey := os.Getenv("XENDIT_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("XENDIT_API_KEY not found in environment variables")
	}
	encodedKey := base64.StdEncoding.EncodeToString([]byte(apiKey))

	// Set headers
	request.Header.Set("Authorization", "Basic "+encodedKey)
	request.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	// Check if response indicates an error
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API request failed with status %d: %v", resp.StatusCode, result)
	}

	return result, nil
}

func SuccessfulAddFund(user_id uint, amount int) error {

	var wallet models.Wallet
	if err := db.GormDB.Table("wallets").Where("user_id = ?", user_id).First(&wallet).Error; err != nil {
		return fmt.Errorf("failed to fetch wallet data: %v", err)
	}

	if err := db.GormDB.
		Table("wallets").
		Where("user_id = ?", user_id).
		Update("balance", wallet.Balance+amount).Error; err != nil {
		return fmt.Errorf("failed to update wallet: %v", err)
	}

	return nil
}

func SuccessfulWithdraw(user_id uint, amount int) error {

	var wallet models.Wallet
	if err := db.GormDB.Table("wallets").Where("user_id = ?", user_id).First(&wallet).Error; err != nil {
		return fmt.Errorf("failed to fetch wallet data: %v", err)
	}

	if err := db.GormDB.
		Table("wallets").
		Where("user_id = ?", user_id).
		Update("balance", wallet.Balance-amount).Error; err != nil {
		return fmt.Errorf("failed to update wallet: %v", err)
	}

	return nil
}
