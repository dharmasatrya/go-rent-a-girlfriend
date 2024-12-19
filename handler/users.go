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
// @Param user body models.User true "User Registration Information"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/register [post]
func UserRegister(c echo.Context) error {
	var req models.User

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
	}

	req.Password = string(hashedPassword)

	var existingUser models.User
	result := db.GormDB.Where("email = ? OR username = ?", req.Email, req.Username).First(&existingUser)
	if result.RowsAffected > 0 {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email or username already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Database error occurred"})
	}

	if err := db.GormDB.Create(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating user"})
	}

	if ok := helper.ActivityLogger(req.ID, "User registered the account"); !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Activity Logger Error"})
	}

	return c.JSON(http.StatusCreated, req)
}

// UserLogin godoc
// @Summary Login user
// @Description Authenticates a user and returns a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login Credentials"
// @Success 200 {object} map[string]string{token=string}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/login [post]
func UserLogin(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
	}

	var user models.User
	if err := db.GormDB.Where("email = ?", req.Email).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password 84"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error Fetching User Data"})
	}

	fmt.Println(user)

	// TODO: Implement login via username as well
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		fmt.Println(user.Password, req.Password, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error generating token"})
	}

	if ok := helper.ActivityLogger(user.ID, "User logged in"); !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Activity Logger Error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func LoginDebug(c echo.Context) error {
	storedHash := "$2a$10$kcTgp0y8qMUQ9NrSy8yKHudJWMmd4m7aV/2jac05G.OVR5/9PYP8m"
	password := "password"

	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Password matches!")
	}
	return c.JSON(http.StatusOK, "ok")
}

// UserCreateGirlProfile godoc
// @Summary Create girl profile
// @Description Creates a profile for a girl user
// @Tags profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param profile body models.Girl true "Girl Profile Information"
// @Success 201 {object} models.Girl
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/profile/girl [post]
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
// @Success 201 {object} models.Boy
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/profile/boy [post]
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
