package chats

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	chatService "github.com/deissh/osu-lazer/api/pkg/services/chat"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

// Join to chat
func Join(c echo.Context) error {
	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	userId, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	channel, err := chatService.Join(uint(userId), uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, channel)
}

// Leave to chat
func Leave(c echo.Context) error {
	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	userId, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	if err = chatService.Leave(uint(userId), uint(channelId)); err != nil {
		return err
	}

	return c.JSON(200, "{}")
}
