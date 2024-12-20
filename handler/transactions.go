package handler

import (
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ShowUserTransactions godoc
// @Summary Get authenticated user's transactions
// @Description Retrieves all transactions where the authenticated user is either the sender or receiver. Shows both incoming and outgoing transactions.
// @Tags transactions
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Transaction "List of user's transactions"
// @Failure 400 {object} map[string]string "User wallet not found"
// @Failure 401 {object} map[string]string "Unauthorized - Invalid or missing token"
// @Failure 500 {object} map[string]string "Server error"
// @Router /users/transactions/history [get]
func ShowUserTransactions(c echo.Context) error {
	var transactions []models.Transaction
	var wallet models.Wallet

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	userID := uint(claims["user_id"].(float64))

	if err := db.GormDB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, "user wallet does not exist")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user wallet")
	}

	if err := db.GormDB.Where("sender_wallet_id = ? OR receiver_wallet_id = ?", wallet.ID, wallet.ID).Find(&transactions).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching transactions")
	}

	return c.JSON(http.StatusOK, transactions)
}

func ShowTransactions(c echo.Context) error {
	var transactions []models.Transaction

	if err := db.GormDB.Find(&transactions).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching transactions")
	}

	return c.JSON(http.StatusOK, transactions)
}

func ShowTransactionById(c echo.Context) error {
	id := c.Param("id")
	var transactions models.Transaction

	if err := db.GormDB.Where("id = ?", id).First(&transactions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, "transaction with that id does not exist")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching transactions")
	}

	return c.JSON(http.StatusOK, transactions)
}
