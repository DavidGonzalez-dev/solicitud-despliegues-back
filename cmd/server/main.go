package main

import (
	"fmt"
	"go-solicitud-despliegues-back/config"
	"go-solicitud-despliegues-back/database"
	"go-solicitud-despliegues-back/database/migrations"
	"go-solicitud-despliegues-back/internal/handler"
	"go-solicitud-despliegues-back/internal/repository"
	"go-solicitud-despliegues-back/internal/routes"
	"go-solicitud-despliegues-back/internal/service"
	"go-solicitud-despliegues-back/internal/usecase"
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

	// Cors config
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Setup database
	db, err := database.NewDatabaseConnection()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	if err := migrations.Migrate(db); err != nil {
		panic(fmt.Sprintf("failed to run database migrations: %v", err))
	}

	// Setup Auth Config
	authConfig := config.NewAuthConfig()
	authenticator, err := config.NewAuthenticator(authConfig)
	if err != nil {
		panic(fmt.Sprintf("failed to create authenticator: %v", err))
	}


	// Services
	azureDevopsService := service.NewAzureDevopsService()

	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Use cases
	userUseCase := usecase.NewUserUseCase(azureDevopsService, userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.NewUserRoutes(e, userHandler, authenticator)



	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
