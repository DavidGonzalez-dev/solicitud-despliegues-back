package handler

import (
	"context"
	"go-solicitud-despliegues-back/internal/usecase"
	customContext "go-solicitud-despliegues-back/pkg/context"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUserOrganizations(c echo.Context)  error
}

type userHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return &userHandler{
		usecase: usecase,
	}
}
	
func (h *userHandler) GetUserOrganizations(c echo.Context) error {


	ctxUser, err := customContext.CurrentUser(c)
	if err != nil {
		return echo.NewHTTPError(401, "failed to get user from context")
	}

	profile, err := h.usecase.GetUserOrganizations(context.WithValue(c.Request().Context(), customContext.UserCtxKey, ctxUser))
	if err != nil {
		return echo.NewHTTPError(500, "failed to get user profile")
	}

	return c.JSON(200, profile)
}
