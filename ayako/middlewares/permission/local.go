package permission

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// MustLogin check
//noinspection GoUnusedGlobalVariable
var MustLogin = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := c.Get("current_user_id")

		_, ok := data.(uint64)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}
