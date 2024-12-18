package handler

import (
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"

	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

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

func XenditCallbackHandler(c echo.Context) error {
	var req models.XenditCallback

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := db.GormDB.
		Table("internal_transactions").
		Where("external_id = ?", req.ExternalId).
		Update("status", req.Status).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update payment status")
	}

	var transaction models.InternalTransaction
	if err := db.GormDB.
		Table("internal_transactions").
		Where("external_id = ?", req.ExternalId).
		First(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch updated transaction")
	}

	if req.Status == "PAID" {
		helper.SuccessfulAddFund(transaction.UserID, transaction.Amount)
	}

	return c.JSON(http.StatusOK, "ok")
}
