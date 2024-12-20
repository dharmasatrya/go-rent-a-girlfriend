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
// @Tags bookings
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

	booking.GirlUserID = req.GirlID
	booking.BookingDate = req.BookingDate
	booking.NumOfDays = req.NumOfDays

	//get userid
	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}
	userId := uint(claims["user_id"].(float64))
	booking.BoyUserID = userId

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
		return echo.NewHTTPError(http.StatusConflict, "Error fetching boy profile")
	}
	booking.Boy = boyProfile

	//calculate total cost
	booking.TotalCost = girlProfile.DailyRate * req.NumOfDays

	//get girl wallet id
	girlWalletId, err := helper.GetWalletIDByUserID(booking.GirlUserID)
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

// GetAllBooking godoc
// @Summary get all booking of a user
// @Description get all booking of a user
// @Tags bookings
// @Accept json
// @Produce json
// @Success 201 {array} models.Booking
// @Failure 500 {object} map[string]string
// @Router /bookings [get]
func GetAllBooking(c echo.Context) error {
	// Initialize our bookings slice
	var bookings []models.Booking

	// Get user ID from JWT token
	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}
	userID := uint(claims["user_id"].(float64))

	// Fetch bookings with related data using Preload
	if err := db.GormDB.
		Preload("Boy").
		Preload("Girl").
		Where("boy_user_id = ?", userID).
		Find(&bookings).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching bookings")
	}

	return c.JSON(http.StatusOK, bookings)
}

// CancelBooking godoc
// @Summary Cancel an existing booking
// @Description Cancels a booking if it belongs to the authenticated user. Only the boy who made the booking can cancel it.
// @Tags bookings
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Booking ID to cancel" example:"1"
// @Success 200 {object} models.Booking "Cancelled booking details"
// @Failure 400 {object} map[string]string "Invalid booking ID"
// @Failure 401 {object} map[string]string "Unauthorized - Invalid or missing token"
// @Failure 403 {object} map[string]string "Forbidden - Not your booking"
// @Failure 404 {object} map[string]string "Booking not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /bookings/{id} [delete]
func CancelBooking(c echo.Context) error {
	bookingID := c.Param("id")
	var booking models.Booking

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}
	userID := uint(claims["user_id"].(float64))

	if err := db.GormDB.
		Preload("Boy").
		Preload("Girl").
		Where("id = ?", bookingID).
		First(&booking).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching bookings")
	}

	if booking.BoyUserID != userID {
		return echo.NewHTTPError(http.StatusForbidden, "Not your booking")
	}

	// Fetch bookings with related data using Preload
	var deletedBooking models.Booking

	if err := db.GormDB.
		Preload("Boy").
		Preload("Girl").
		Where("id = ?", bookingID).
		Delete(&deletedBooking).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching bookings")
	}

	return c.JSON(http.StatusOK, booking)
}
