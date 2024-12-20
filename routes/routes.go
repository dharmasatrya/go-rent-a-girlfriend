package routes

import (
	"rent-a-girlfriend/controller"
	"rent-a-girlfriend/handler"
	repository "rent-a-girlfriend/repositories"
	"rent-a-girlfriend/service"

	m "rent-a-girlfriend/middleware"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {

	gfr := repository.NewGirlfriendsRepository(db)
	gs := service.NewGirlfriendService(gfr)
	gfc := controller.NewGirlfriendController(gs)

	u := e.Group("/users")
	u.POST("/register", handler.UserRegister)
	u.POST("/login", handler.UserLogin)
	up := u.Group("/profiles")
	up.Use(echojwt.JWT([]byte("secret")))
	up.POST("/girls", handler.UserCreateGirlProfile, m.RequireRole("girls"))
	up.POST("/boys", handler.UserCreateBoyProfile, m.RequireRole("boys"))
	th := u.Group("/transactions")
	th.Use(echojwt.JWT([]byte("secret")))
	th.GET("/history", handler.ShowUserTransactions)

	g := e.Group("/girlfriends")
	g.Use(echojwt.JWT([]byte("secret")))
	g.Use(m.RequireRole("boys"))
	g.GET("", handler.GetAvailableGirls)
	g.GET("/:id", gfc.GetGirlById)
	g.POST("/ratings", handler.GiveRating)

	b := e.Group("/bookings")
	b.Use(echojwt.JWT([]byte("secret")))
	b.Use(m.RequireRole("boys"))
	b.POST("", handler.CreateBooking)
	b.GET("", handler.GetAllBooking)
	b.DELETE("/:id", handler.CancelBooking)

	w := e.Group("/wallets")
	w.Use(echojwt.JWT([]byte("secret")))
	w.POST("", handler.CreateWallet)
	w.POST("/withdrawal", handler.WithdrawFunds)
	w.POST("/deposit", handler.DepositFunds, m.RequireRole("boys"))

	adm := e.Group("/admin")
	adm.Use(m.RequireRole("admin"))
	adm.GET("/transactions", handler.ShowTransactions)
	adm.GET("/transactions/:id", handler.ShowTransactionById)

	e.POST("/xenditcallback/invoice", handler.XenditInvoiceCallbackHandler)
	e.POST("/xenditcallback/disbursement", handler.XenditDisbursementCallbackHandler)
}
