package handler

import (
	"fmt"
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/models"
	repository "rent-a-girlfriend/repositories"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type GirlfriendHandler struct {
	gfRepo repository.GirlfriendsRepository
}

func NewGirlfriendHandler(gfRepo repository.GirlfriendsRepository) *GirlfriendHandler {
	return &GirlfriendHandler{
		gfRepo: gfRepo,
	}
}

// GetAvailableGirls godoc
// @Summary Get available girls with optional filtering
// @Description Retrieves a list of available girls. Can be filtered by date and age. If no filters are provided, returns all available girls.
// @Tags girlfriends
// @Accept json
// @Produce json
// @Param date query string false "Date to check availability (format: 2024-01-01)" example:"2024-01-01"
// @Param age query string false "Filter by specific age" example:"23"
// @Success 200 {array} models.Girl "List of available girls"
// @Failure 400 {object} map[string]string "Invalid age parameter"
// @Failure 500 {object} map[string]string "Server error"
// @Router /girlfriends [get]
func GetAvailableGirls(c echo.Context) error {
	date := c.QueryParam("date")
	age := c.QueryParam("age")

	query := db.GormDB.
		Joins("LEFT JOIN availabilities ON girls.id = availabilities.girl_id")

	conditions := make([]string, 0)
	values := make([]interface{}, 0)

	if date == "" {
		conditions = append(conditions, "availabilities.is_available IS TRUE")
	} else {
		conditions = append(conditions, "(availabilities.is_available IS TRUE OR ? NOT BETWEEN availabilities.start_date AND availabilities.end_date)")
		values = append(values, date)
	}

	if age != "" {
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid age parameter")
		}

		conditions = append(conditions, "girls.age = ?")
		values = append(values, ageInt)
	}

	whereClause := strings.Join(conditions, " AND ")

	var girls []models.Girl
	if err := query.Where(whereClause, values...).Find(&girls).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching availabilities")
	}

	if girls == nil {
		girls = make([]models.Girl, 0)
	}

	return c.JSON(http.StatusOK, girls)
}

// GetGirlById godoc
// @Summary Get a specific girl's profile by ID
// @Description Retrieves detailed information about a girl's profile using their ID
// @Tags girlfriends
// @Accept json
// @Produce json
// @Param id path int true "Girl ID" example:"1"
// @Success 200 {object} models.GirlDetailResponse "Girl profile details"
// @Failure 404 {object} map[string]string "Girl not found"
// @Failure 500 {object} map[string]string "Server error"
// @Router /girlfriends/{id} [get]
// func (h *GirlfriendHandler) GetGirlById(c echo.Context) error {
// 	girlIdStr := c.Param("id")
// girlId, err := strconv.Atoi(girlIdStr)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "error converting string")
// 	}
// 	// var girl models.Girl
// 	// var availability models.Availability
// 	// var ratings []models.Rating

// 	girlData, err := h.gfRepo.GetGirlfriendById(girlId)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
// 	}
// 	// if err := h.db.Where("id = ?", girlId).First(&girl).Error; err != nil {
// 	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 	// 		return echo.NewHTTPError(http.StatusNotFound, "Girl not found")
// 	// 	}
// 	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching girl profile")
// 	// }

// 	// if err := h.db.Where("girl_id = ?", girlId).First(&availability).Error; err != nil {
// 	// 	availability.IsAvailable = false
// 	// }

// 	// if err := h.db.Where("girl_id = ?", girlId).Find(&ratings).Error; err != nil {
// 	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching ratings")
// 	// }

// 	// response := models.GirlDetailResponse{
// 	// 	ID:                girl.ID,
// 	// 	UserID:            girl.UserID,
// 	// 	FirstName:         girl.FirstName,
// 	// 	LastName:          girl.LastName,
// 	// 	Age:               girl.Age,
// 	// 	ProfilePictureURL: girl.ProfilePictureURL,
// 	// 	Bio:               girl.Bio,
// 	// 	DailyRate:         girl.DailyRate,
// 	// 	IsAvailable:       availability.IsAvailable,
// 	// 	Ratings:           ratings,
// 	// }

// 	return c.JSON(http.StatusOK, girlData)
// }

// GiveRating godoc
// @Summary Submit a rating and review for a girl
// @Description Allows users to submit ratings (1-5 stars) and written reviews for girls they have booked
// @Tags girlfriends
// @Accept json
// @Produce json
// @Security Bearer
// @Param rating body models.Rating true "Rating and review details"
// @Success 201 {object} models.GiveRatingRequest "Created rating details"
// @Failure 400 {object} map[string]string "Invalid request payload or rating value"
// @Failure 401 {object} map[string]string "Unauthorized - Invalid or missing token"
// @Failure 403 {object} map[string]string "Forbidden - Can't rate without a prior booking"
// @Failure 500 {object} map[string]string "Server error"
// @Router /girlfriends/ratings [post]
func GiveRating(c echo.Context) error {
	var review models.Rating

	if err := c.Bind(&review); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := db.GormDB.Debug().Create(&review).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating review")
	}

	return c.JSON(http.StatusOK, review)
}
