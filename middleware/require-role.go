package customMiddleware

import (
	customContext "go-solicitud-despliegues-back/pkg/context"
	pkgHttp "go-solicitud-despliegues-back/pkg/http"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

func RequireRole(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get(string(customContext.UserCtxKey))
			if user == nil {
				return c.JSON(http.StatusForbidden, pkgHttp.HttpError{
					Status:  http.StatusForbidden,
					Message: "Access denied: no user in context",
					Error:   "no user in context",
				})
			}

			ctxUser, ok := user.(customContext.ContextUser)
			if !ok {
				return c.JSON(http.StatusForbidden, pkgHttp.HttpError{
					Status:  http.StatusForbidden,
					Message: "Access denied: invalid user in context",
					Error:   "invalid user in context",
				})
			}

			if slices.Contains(roles, ctxUser.Role) {
				return next(c)
			}
			return c.JSON(http.StatusForbidden, pkgHttp.HttpError{
				Status:  http.StatusForbidden,
				Message: "Access denied: insufficient permissions",
				Error:   "insufficient permissions",
			})
		}
	}
}
