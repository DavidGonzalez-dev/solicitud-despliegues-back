package main

import (
	"fmt"
	"go-solicitud-despliegues-back/internal/config"
	"go-solicitud-despliegues-back/internal/handler"
	"go-solicitud-despliegues-back/internal/repository/migrations"
	"go-solicitud-despliegues-back/internal/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Load env variables
	if err := godotenv.Load(); err != nil {
		panic("failed to load env variables")
	}

	// Do database migrations.

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Connect to database
	db, err := config.NewDatabaseConnection()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Run database migrations.
	if err := migrations.Migrate(db); err != nil {
		panic(fmt.Sprintf("failed to run database migrations: %v", err))
	}

	// Wiring dependencies
	azureDevopsService := service.NewAzureDevopsService()
	userHandler := handler.NewUserHandler(azureDevopsService)

	e.POST("/me", userHandler.GetUserAzureDVProfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
