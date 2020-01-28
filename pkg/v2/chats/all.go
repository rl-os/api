package chats

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	chatService "github.com/deissh/osu-api-server/pkg/services/chat"
	"github.com/labstack/echo/v4"
)

// GetAll of all joinable public channels
func GetAll(c echo.Context) error {
	channels, err := chatService.GetChannels()
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

// GetAll of all user channels
func GetJoinedAll(c echo.Context) error {
	data := c.Get("current_user")

	current, ok := data.(*entity.User)
	if !ok {
		return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
	}

	channels, err := chatService.GetUserChannels(current.ID)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}
