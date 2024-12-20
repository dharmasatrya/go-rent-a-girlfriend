package middleware

import (
	"fmt"
	"net/http"
	"rent-a-girlfriend/helper"

	"github.com/labstack/echo/v4"
)

func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := helper.GetClaimsFromToken(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			userRole := claims["user_role"].(string)
			if userRole != role {
				fmt.Println(userRole)
				return echo.NewHTTPError(http.StatusForbidden, "Access denied")
			}

			return next(c)
		}
	}
}
