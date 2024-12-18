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

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Rent A Girlfriend API
// @version 1.0
// @description This is a girlfriend rental service API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	db.InitDB()

	defer func() {
		log.Println("Closing database connection...")
		db.CloseDB()
	}()

	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
