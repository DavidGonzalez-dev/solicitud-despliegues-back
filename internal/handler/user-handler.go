package handler

import (
	"context"
	"go-solicitud-despliegues-back/internal/usecase"
	customContext "go-solicitud-despliegues-back/pkg/context"
	pkgHttp "go-solicitud-despliegues-back/pkg/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUserInfo(c echo.Context)  error
}

type userHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return &userHandler{
		usecase: usecase,
	}
}
	
func (h *userHandler) GetUserInfo(c echo.Context) error {


	ctxUser, err := customContext.CurrentUser(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, pkgHttp.HttpError{
			Status: http.StatusNotFound,
			Message: "Error while getting user from context",
			Error: err.Error(),
		})
	}

	profile, err := h.usecase.GetUserInfo(context.WithValue(c.Request().Context(), customContext.UserCtxKey, ctxUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkgHttp.HttpError{
			Status: http.StatusInternalServerError,
			Message: "Failed to get user profile",
			Error: err.Error(),
		})
	}

	return c.JSON(200, profile)
}
