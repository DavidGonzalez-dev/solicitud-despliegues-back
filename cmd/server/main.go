package main

import (
	"go-solicitud-despliegues-back/internal/handler"
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

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Wiring dependencies
	oboService := service.NewOboService()
	oboUseCase := usecase.NewOboUsecase(oboService)
	oboHandler := handler.NewOboHandler(oboUseCase)

	e.POST("/obo", oboHandler.LoginOnBehalfOf)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
