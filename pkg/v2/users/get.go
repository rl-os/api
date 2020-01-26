package users

import (
	"github.com/deissh/osu-api-server/pkg"
	oauthService "github.com/deissh/osu-api-server/pkg/services/oauth"
	userService "github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/deissh/osu-api-server/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

var modes = []string{"std", "mania", "catch", "taiko"}

// GetUserByID handler
func GetUserByID(c echo.Context) (err error) {
	mode := c.Param("mode")
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	userID, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err := userService.GetUser(uint(userID), mode)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// GetUserByToken handler
func GetUserByToken(c echo.Context) error {
	mode := c.Param("mode")
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	token, ok := c.Get("current_user_token").(oauthService.Token)
	log.Debug().Interface("token", token).Send()
	if ok != true {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err := userService.GetUser(token.UserID, mode)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
