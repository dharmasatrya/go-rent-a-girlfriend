package handler

import (
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"

	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// CreateWallet godoc
// @Summary Create a wallet
// @Description Registers a new wallet
// @Tags wallet
// @Accept json
// @Produce json
// @Param user body models.CreateWalletRequest true "User wallet"
// @Success 201 {object} models.Wallet
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /wallets [post]
func CreateWallet(c echo.Context) error {
	var req models.Wallet

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	req.UserID = uint(claims["user_id"].(float64))
	req.Balance = 0

	var existingWallet models.Wallet
	result := db.GormDB.Where("user_id = ?", req.UserID).First(&existingWallet)
	if result.RowsAffected > 0 {
		return echo.NewHTTPError(http.StatusConflict, "User already has a wallet")
	}

	if err := db.GormDB.Create(&req).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating wallet")
	}

	return c.JSON(http.StatusCreated, req)
}

// DepositFunds godoc
// @Summary Create a deposit
// @Description Create a deposit
// @Tags wallet
// @Accept json
// @Produce json
// @Param user body models.DepostitRequest true "Deposit amount"
// @Success 200 {object} models.Wallet
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /wallets/deposit [post]
func DepositFunds(c echo.Context) error {
	var req models.DepostitRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	userID := uint(claims["user_id"].(float64))
	externalId, err := gonanoid.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating invoice UUID")
	}

	var user models.User
	if err := db.GormDB.Where("id = ?", userID).Take(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching userdata")
	}

	var boyProfile models.Boy
	if err := db.GormDB.Where("user_id = ?", userID).First(&boyProfile).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl profile")
	}

	invoiceReq := models.XenditInvoiceRequest{
		ExternalId:  externalId,
		Amount:      req.Amount,
		Description: "Booking Payment",
		FirstName:   boyProfile.FirstName,
		LastName:    boyProfile.LastName,
		Email:       user.Email,
		Phone:       "+6281299640904",
	}

	InternalTransaction := models.InternalTransaction{
		UserID:     userID,
		ExternalId: externalId,
		Amount:     req.Amount,
		Status:     "UNPAID",
		Type:       "MONEY_IN",
	}

	if err := db.GormDB.Create(&InternalTransaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating profile")
	}

	response, err := helper.CreateXenditInvoice(invoiceReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create payment: "+err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

// CreateWallet godoc
// @Summary Create a withdrawal
// @Description Create a withdrawal
// @Tags wallet
// @Accept json
// @Produce json
// @Param user body models.WithdrawalRequest true "withdrawal"
// @Success 200 {object} models.Wallet
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /wallets/withdrawal [post]
func WithdrawFunds(c echo.Context) error {
	var req models.WithdrawalRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	userID := uint(claims["user_id"].(float64))
	externalId, err := gonanoid.New()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating invoice UUID")
	}

	var user models.User
	if err := db.GormDB.Where("id = ?", userID).Take(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching userdata")
	}

	var wallet models.Wallet
	if err := db.GormDB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching wallet details")
	}

	if wallet.Balance < req.Amount {
		return echo.NewHTTPError(http.StatusBadRequest, "Insufficient balance")
	}

	disbursementRequest := models.XenditDisbursementRequest{
		ExternalId:        externalId,
		Amount:            req.Amount,
		Description:       "Withdrawal",
		BankCode:          "BCA",
		AccountHolderName: wallet.BankAccountName,
		BankAccountNumber: wallet.BankAccountNumber,
		Email:             user.Email,
	}

	InternalTransaction := models.InternalTransaction{
		UserID:     userID,
		ExternalId: externalId,
		Amount:     req.Amount,
		Status:     "PENDING",
		Type:       "MONEY_OUT",
	}

	if err := db.GormDB.Create(&InternalTransaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating profile")
	}

	response, err := helper.CreateXenditDisbursement(disbursementRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create payment: "+err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
