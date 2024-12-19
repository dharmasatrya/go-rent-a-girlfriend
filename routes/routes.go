package routes

import (
	"rent-a-girlfriend/handler"

	m "rent-a-girlfriend/middleware"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	u := e.Group("/users")
	u.POST("/register", handler.UserRegister)
	u.POST("/login", handler.UserLogin)
	up := u.Group("/profiles")
	up.Use(echojwt.JWT([]byte("secret")))
	up.POST("/girls", handler.UserCreateGirlProfile, m.RequireRole("girls"))
	up.POST("/boys", handler.UserCreateBoyProfile, m.RequireRole("boys"))

	g := e.Group("/girlfriends")
	g.Use(echojwt.JWT([]byte("secret")))
	g.Use(m.RequireRole("boys"))
	g.GET("", handler.GetAvailableGirls)
	g.GET("/:id", handler.GetGirlById)
	g.POST("/review", handler.GiveRating)

	b := e.Group("/bookings")
	b.Use(echojwt.JWT([]byte("secret")))
	b.Use(m.RequireRole("boys"))
	b.POST("", handler.CreateBooking)
	b.GET("", handler.GetAllBooking)
	// b.DELETE("/:id", )

	w := e.Group("/wallets")
	w.Use(echojwt.JWT([]byte("secret")))
	w.POST("", handler.CreateWallet)
	w.POST("/withdrawal", handler.WithdrawFunds)
	w.POST("/deposit", handler.DepositFunds, m.RequireRole("admin"))

	adm := e.Group("/admin")
	adm.Use(m.RequireRole("admin"))
	// adm.GET("/bookings")

	e.POST("/xenditcallback/invoice", handler.XenditInvoiceCallbackHandler)
	e.POST("/xenditcallback/disbursement", handler.XenditDisbursementCallbackHandler)
}
