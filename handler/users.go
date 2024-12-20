package handler

import (
	"fmt"
	"net/http"
	"rent-a-girlfriend/db"
	"rent-a-girlfriend/helper"
	"rent-a-girlfriend/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("secret")

// UserRegister godoc
// @Summary Register a new user
// @Description Registers a new user with username, email, password and role
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "User Registration Information"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/register [post]
func UserRegister(c echo.Context) error {
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
	}

	req.Password = string(hashedPassword)

	fmt.Printf("Stored hash in DB: %s\n", hashedPassword)

	var existingUser models.User
	result := db.GormDB.Table("users").Where("email = ? OR username = ?", req.Email, req.Username).First(&existingUser)
	if result.RowsAffected > 0 {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email or username already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Database error occurred"})
	}

	createdUser := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
	}

	if err := db.GormDB.Table("users").Create(&createdUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating user"})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// UserLogin godoc
// @Summary Login user
// @Description Authenticates a user and returns a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequestWithPassword true "Login Credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/login [post]
func UserLogin(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
	}

	var user models.User
	if err := db.GormDB.Debug().Table("users").Where("email = ?", req.Email).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error Fetching User Data"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Printf("Password comparison error: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
	}

	fmt.Println(user.Role)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error generating token"})
	}

	return c.JSON(http.StatusOK, models.LoginResponse{Token: tokenString})
}

// UserCreateGirlProfile godoc
// @Summary Create girl profile
// @Description Creates a profile for a girl user
// @Tags profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param profile body models.CreateGirlProfileRequest true "Girl Profile Information"
// @Success 201 {object} models.Girl
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/profiles/girls [post]
func UserCreateGirlProfile(c echo.Context) error {
	var req models.Girl

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	req.UserID = uint(claims["user_id"].(float64))

	var existingProfile models.Girl
	result := db.GormDB.Where("user_id = ?", req.UserID).First(&existingProfile)
	if result.RowsAffected > 0 {
		return echo.NewHTTPError(http.StatusConflict, "User already has a girl profile")
	}

	if err := db.GormDB.Create(&req).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating profile")
	}

	return c.JSON(http.StatusCreated, req)
}

// UserCreateBoyProfile godoc
// @Summary Create boy profile
// @Description Creates a profile for a boy user
// @Tags profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param profile body models.Boy true "Boy Profile Information"
// @Success 201 {object} models.CreateBoyProfileRequest
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/profiles/boys [post]
func UserCreateBoyProfile(c echo.Context) error {
	var req models.Boy

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching claims from token")
	}

	req.UserID = uint(claims["user_id"].(float64))

	var existingProfile models.Boy
	result := db.GormDB.Where("user_id = ?", req.UserID).First(&existingProfile)
	if result.RowsAffected > 0 {
		return echo.NewHTTPError(http.StatusConflict, "User already has a boy profile")
	}

	if err := db.GormDB.Create(&req).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating profile")
	}

	return c.JSON(http.StatusCreated, req)
}
