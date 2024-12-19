package handler

import (
	"fmt"
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"
	"strings"

	"github.com/labstack/echo/v4"
)

// CreateBooking godoc
// @Summary Book a new date
// @Description Book a date
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body models.BookingRequest true "Booking Information"
// @Success 201 {object} models.Booking
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /bookings [post]
func CreateBooking(c echo.Context) error {
	var req models.BookingRequest
	var booking models.Booking

	//bind req
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	booking.GirlID = req.GirlID
	booking.BookingDate = req.BookingDate
	booking.NumOfDays = req.NumOfDays

	//get userid
	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}
	userId := uint(claims["user_id"].(float64))

	//get girl availability
	var availability models.Availability
	if err := db.GormDB.Where("girl_id = ? AND is_available = ?", req.GirlID, true).First(&availability).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Girl is not available")
	}

	//get girl profile
	var girlProfile models.Girl
	if err := db.GormDB.Where("id = ?", req.GirlID).First(&girlProfile).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl profile")
	}
	booking.Girl = girlProfile

	//get boy profile
	var boyProfile models.Boy
	if err := db.GormDB.Where("user_id = ?", userId).First(&boyProfile).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl profile")
	}
	booking.Boy = boyProfile

	//calculate total cost
	booking.TotalCost = girlProfile.DailyRate * req.NumOfDays

	//get girl wallet id
	girlWalletId, err := helper.GetWalletIDByUserID(booking.GirlID)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl wallet")
	}
	//get boy wallet id
	boyWalletId, err := helper.GetWalletIDByUserID(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching boy wallet")
	}

	//transaction
	if err := helper.Transaction(boyWalletId, girlWalletId, booking.TotalCost); err != nil {
		// Handle error (e.g., insufficient balance, database errors)
		if strings.Contains(err.Error(), "insufficient balance") {
			// Handle insufficient balance specifically
			return echo.NewHTTPError(http.StatusBadRequest, "insufficient balance")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Transaction failed")
	}

	// Create the booking
	fmt.Println(booking)
	if err := db.GormDB.Debug().Create(&booking).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating booking")
	}

	// Update availability
	endDate := req.BookingDate.AddDate(0, 0, req.NumOfDays)
	if err := db.GormDB.Table("availabilities").
		Where("girl_id = ?", req.GirlID).
		Updates(map[string]interface{}{
			"is_available": false,
			"start_date":   req.BookingDate,
			"end_date":     endDate,
		}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating availability")
	}

	return c.JSON(http.StatusCreated, booking)
}

func GetAllBooking(c echo.Context) error {
	var booking []models.Booking

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}
	boyID := uint(claims["user_id"].(float64))

	if err := db.GormDB.Where("boy_id = ?", boyID).Find(&booking).Error; err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Error fetching girl profile")
	}

	fmt.Println(boyID)

	return c.JSON(http.StatusOK, booking)
}

func GetAvailableGirls(c echo.Context) error {
	date := c.QueryParam("date")
	var girls []models.Girl

	// Get girls who don't have bookings for the specified date
	if err := db.GormDB.
		Joins("LEFT JOIN availabilities ON girls.id = availabilities.girl_id").
		Where("availabilities.id IS NULL OR ? NOT BETWEEN availabilities.start_date AND availabilities.end_date", date).
		Find(&girls).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching availabilities")
	}

	return c.JSON(http.StatusOK, girls)
}
