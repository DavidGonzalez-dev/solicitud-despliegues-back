package customContext

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// const to store user context key
const UserCtxKey string = "auth.user"

// Context user contains the min info we want to store in the context after validating the JWT
type ContextUser struct {
	AccessToken string        // access token sent
	Subject     string        // sub
	OID         string        // oid (object id)
	UPD         string        // upd (preferred username)
	Scopes      []string      // scp (scopes)
	Roles       []any         // roles (roles)
	Raw         jwt.MapClaims // raw claims (all claims)
}

// Func to get current user from context
func CurrentUser(c echo.Context) (ContextUser, error) {

	value := c.Get(string(UserCtxKey))
	if value == nil {
		return ContextUser{}, errors.New("no user in context")
	}

	userContext, ok := value.(ContextUser)
	if !ok {
		return ContextUser{}, errors.New("invalid user in context")
	}

	return userContext, nil
}
