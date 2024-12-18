package helper

import (
	"fmt"
	"net/http"
	"regexp"
	"rent-a-girlfriend/db"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

type UserActivityLog struct {
	ID          uint
	UserId      uint
	Description string
}

func GetClaimsFromToken(c echo.Context) (jwt.MapClaims, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		fmt.Println(ok)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Error Fetching Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Token Claims Error")
	}

	return claims, nil
}

func IsValidURL(url string) bool {
	regex := `^((https?|ftp):\/\/)?[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\/[^\s]*)?\.(jpg|jpeg|png|gif|bmp|svg|webp)$`
	re := regexp.MustCompile(regex)

	return re.MatchString(url)
}

func ActivityLogger(userId uint, description string) bool {
	log := UserActivityLog{
		UserId:      userId,
		Description: description,
	}

	if err := db.GormDB.Debug().Create(&log).Error; err != nil {
		return false
	}

	return true
}