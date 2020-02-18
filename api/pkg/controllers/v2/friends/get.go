package friends

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	userService "github.com/deissh/osu-lazer/api/pkg/services/user"
	"github.com/labstack/echo/v4"
)

// Get user friends handler
func Get(c echo.Context) (err error) {
	data := c.Get("current_user")

	current, ok := data.(*entity.User)
	if !ok {
		return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
	}

	users, err := userService.GetSubscriptions(current.ID)
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
