package handler

import (
	"fmt"
	"go-solicitud-despliegues-back/internal/domain"
	"go-solicitud-despliegues-back/internal/usecase"
	"net/http"

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
	var req domain.OboRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	fmt.Println(req)

	res, err := h.OboUseCase.LoginOnBehalfOf(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error", "error_message": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
