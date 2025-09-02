package routes

import (
	"go-solicitud-despliegues-back/config"
	"go-solicitud-despliegues-back/internal/handler"
	customMiddleware "go-solicitud-despliegues-back/middleware"

	"github.com/labstack/echo/v4"
)


func NewUserRoutes(e *echo.Echo, userHandler handler.UserHandler, authenticator *config.Authenticator) {

	userGroup := e.Group("/user")

	userGroup.POST("/me", userHandler.GetUserOrganizations, customMiddleware.RequireAccessToken(authenticator))
}
