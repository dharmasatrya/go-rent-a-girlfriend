package routes

import (
	"rent-a-girlfriend/handler"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	u := e.Group("/users")
	u.POST("/register", handler.UserRegister)
	u.POST("/login", handler.UserLogin)
	up := u.Group("/profiles")
	up.Use(echojwt.JWT([]byte("secret")))
	up.POST("/girls", handler.UserCreateGirlProfile)
	up.POST("/boys", handler.UserCreateBoyProfile)

	// g := e.Group("/girlfirends")
	// g.Use(echojwt.JWT([]byte("secret")))
	// g.GET("", get all available girlfriend)
	// g.GET("/:id", get by id)

	b := e.Group("/bookings")
	b.Use(echojwt.JWT([]byte("secret")))
	b.POST("", handler.CreateBooking)
	// b.GET("", get user booking)
	// b.DELETE("/:id", cancel a booking)

	w := e.Group("/wallets")
	w.Use(echojwt.JWT([]byte("secret")))
	// w.POST("/withdrawal", handler.)
	w.POST("/deposit", handler.DepositFunds)

	e.POST("/xenditcallback", handler.XenditCallbackHandler)
}
