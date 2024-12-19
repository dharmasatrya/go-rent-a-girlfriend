package handler

import (
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/models"

	"github.com/labstack/echo/v4"
)

func GiveRating(c echo.Context) error {
	var review models.Rating

	if err := c.Bind(&review); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := db.GormDB.Debug().Create(&review).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating review")
	}

	return c.JSON(http.StatusOK, review)
}
