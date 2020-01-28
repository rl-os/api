package chats

import (
	"github.com/deissh/osu-api-server/pkg"
	chatService "github.com/deissh/osu-api-server/pkg/services/chat"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetAllMessages in channel
func GetAllMessages(c echo.Context) error {
	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	messages, err := chatService.GetMessages(uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}
