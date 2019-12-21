package users

import (
	"github.com/deissh/osu-api-server/pkg"
	oauthService "github.com/deissh/osu-api-server/pkg/services/oauth"
	userService "github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

// GetUserByID handler
func GetUserByID(c echo.Context) (err error) {
	var user userService.DetailedUser

	userID, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err = userService.GetDetailedUser(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// GetUserByToken handler
func GetUserByToken(c echo.Context) (err error) {
	var user userService.DetailedUser

	token, ok := c.Get("uset_token").(oauthService.Token)
	log.Debug().Interface("token", token).Send()
	if ok != true {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err = userService.GetDetailedUser(token.UserID)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
