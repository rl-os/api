package scopes

import (
	"fmt"
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/common/utils"
	"github.com/deissh/osu-lazer/api/pkg/services/oauth"
	"github.com/labstack/echo/v4"
	"strings"
)

// Required scopes
func Required(required ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			data := c.Get("current_user_token")

			token, ok := data.(*oauth.Token)
			if !ok {
				return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
			}

			if token.Scopes == "*" || len(required) == 0 {
				return next(c)
			}

			scopes := strings.SplitN(token.Scopes, ",", -1)

			if utils.ContainsStrings(scopes, required) != true {
				return pkg.NewHTTPError(401, "oauth", fmt.Sprint("Required %a scopes, but received %a", required, scopes))
			}

			return next(c)
		}
	}
}
