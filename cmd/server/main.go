package main

import (
	"go-solicitud-despliegues-back/internal/handler"
	"go-solicitud-despliegues-back/internal/service"
	"go-solicitud-despliegues-back/internal/usecase"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Wiring dependencies
	oboService := service.NewOboService("user.read")
	oboUseCase := usecase.NewOboUsecase(oboService)
	oboHandler := handler.NewOboHandler(oboUseCase)

	e.POST("/obo", oboHandler.LoginOnBehalfOf)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
