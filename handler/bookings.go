package handler

import (
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateBooking(c echo.Context) error {
	var req models.Booking

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	req.BoyID = uint(claims["user_id"].(float64))

	var girlProfile models.Girl
	if err := db.GormDB.Where("user_id = ?", req.GirlID).First(&girlProfile).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl profile")
	}

	var boyProfile models.Boy
	if err := db.GormDB.Where("user_id = ?", req.BoyID).First(&boyProfile).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl profile")
	}

	req.Boy = boyProfile
	req.Girl = girlProfile

	req.TotalCost = girlProfile.DailyRate * req.NumOfDays

	girlWalletId, err := helper.GetWalletIDByUserID(req.GirlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl wallet")
	}

	boyWalletId, err := helper.GetWalletIDByUserID(req.BoyID)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl wallet")
	}

	if err := helper.Transaction(boyWalletId, girlWalletId, req.TotalCost); err != nil {
		// Handle error (e.g., insufficient balance, database errors)
		if strings.Contains(err.Error(), "insufficient balance") {
			// Handle insufficient balance specifically
			return echo.NewHTTPError(http.StatusBadRequest, "insufficient balance")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Transaction failed")
	}

	// Create the booking first
	if err := db.GormDB.Create(&req).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating booking")
	}

	return c.JSON(http.StatusCreated, req)
}
