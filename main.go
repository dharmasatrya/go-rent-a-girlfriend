package main

import (
	"log"
	"os"
	"os/signal"
	"rent-a-girlfriend/db"
	m "rent-a-girlfriend/middleware"
	"rent-a-girlfriend/routes"
	"syscall"

	_ "rent-a-girlfriend/docs" // Import the generated Swagger docs

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title       My API
// @version     1.0
// @description This is an example API that provides user registration, login, and post management
// @contact.name Developer Name
// @contact.url http://example.com
// @contact.email developer@example.com
// @host
func main() {
	db.InitDB()

	defer func() {
		log.Println("Closing database connection...")
		db.CloseDB()
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	e.HTTPErrorHandler = m.ErrorHandlerMiddleware

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.Init(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := e.Start(":" + port); err != nil {
			if err.Error() != "http: Server closed" {
				log.Fatalf("Error starting server: %v", err)
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down server...")
	if err := e.Shutdown(nil); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server gracefully stopped.")
}
