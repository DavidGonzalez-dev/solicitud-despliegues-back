package customMiddleware

import (
	"go-solicitud-despliegues-back/config"
	customContext "go-solicitud-despliegues-back/pkg/context"
	pkgHttp "go-solicitud-despliegues-back/pkg/http"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
)

// Middleware returns the echo.MiddlewareFunc that validates the token and puts the user in the context.
// Validations performed:
// - Authorization header present and with Bearer
// - valid signature against JWKS
// - issuer == expected
// - audience contains ApiAudience
// - exp (expiration) not expired
// - exposes useful claims in ContextUser
func RequireAccessToken(authenticator *config.Authenticator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Get token from header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, pkgHttp.HttpError{
					Status:  http.StatusUnauthorized,
					Message: "Missing or invalid Authorization header",
					Error:   "Unauthorized",
				})
			}
			tokenString := authHeader[len("Bearer "):]

			// Parse and validate
			token, err := jwt.Parse(tokenString, authenticator.JWKS.Keyfunc)
			if err != nil {
				println("Error parsing token:", err.Error())
				return c.JSON(http.StatusUnauthorized, pkgHttp.HttpError{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})
			}

			// Read Claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				println("Invalid token claims")
				return c.JSON(http.StatusUnauthorized, pkgHttp.HttpError{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})

			}

			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					println("Token has expired")
					return c.JSON(http.StatusUnauthorized, pkgHttp.HttpError{
						Status:  http.StatusUnauthorized,
						Message: "Unauthorized",
					})
				}
			}

			// Validate claims
			if claims["aud"] != authenticator.Audience {
				println("Invalid audience")
				return c.JSON(http.StatusUnauthorized, pkgHttp.HttpError{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})
			}
			if claims["iss"] != authenticator.Issuer {
				println("Invalid issuer")
				return c.JSON(http.StatusUnauthorized, pkgHttp.HttpError{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})
			}

			// Store Map claims as contextUser
			contextUser := customContext.ContextUser{
				AccessToken: tokenString,
				Subject:     claims["sub"].(string),
				OID:         claims["oid"].(string),
				Roles:       claims["roles"].([]any),
				Raw:         claims,
			}
			c.Set(customContext.UserCtxKey, contextUser)
			return next(c)

		}
	}
}
