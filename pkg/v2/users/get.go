package users

import (
	"github.com/deissh/osu-api-server/pkg"
	userService "github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUserById handler
func GetUserById(c echo.Context) (err error) {
	var user userService.DetailedUser

	userId, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err = userService.GetDetailedUser(uint(userId))
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
