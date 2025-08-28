package handler

import (
	"go-solicitud-despliegues-back/internal/domain"
	"go-solicitud-despliegues-back/internal/usecase"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type OboHandler interface {
	LoginOnBehalfOf(c echo.Context) error
}

type oboHandler struct {
	OboUseCase usecase.OboUseCase
}

func NewOboHandler(oboUseCase usecase.OboUseCase) OboHandler {
	return &oboHandler{
		OboUseCase: oboUseCase,
	}
}

func (h *oboHandler) LoginOnBehalfOf(c echo.Context) error {

	// Get token from request
	token, ok := c.Request().Header["Authorization"]
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
	}

	tokenString := strings.Join(token, "")
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	req := domain.OboRequest{
		AccessToken: tokenString,
	}

	// Get OBO token
	res, err := h.OboUseCase.LoginOnBehalfOf(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error", "error_message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
