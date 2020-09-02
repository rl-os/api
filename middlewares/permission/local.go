package permission

import (
	"context"
	"github.com/labstack/echo/v4"
	myctx "github.com/rl-os/api/ctx"
	"net/http"
)

// MustLogin check
//noinspection GoUnusedGlobalVariable
var MustLogin = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, ok := c.Get("context").(context.Context)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Invalid request context")
		}

		_, err := myctx.GetUserToken(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}
