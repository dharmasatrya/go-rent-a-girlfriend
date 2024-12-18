package handler

import (
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"

	"github.com/labstack/echo/v4"
)

func XenditInvoiceCallbackHandler(c echo.Context) error {
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

func XenditDisbursementCallbackHandler(c echo.Context) error {
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

	if req.Status == "COMPLETED" {
		helper.SuccessfulWithdraw(transaction.UserID, transaction.Amount)
	}

	return c.JSON(http.StatusOK, "ok")
}
