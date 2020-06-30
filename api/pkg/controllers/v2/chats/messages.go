package chats

import (
	"github.com/deissh/rl/api/pkg"
	chatService "github.com/deissh/rl/api/pkg/services/chat"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetUpdatesRequestData contain incoming data
type GetMessagesRequestData struct {
	Limit uint `json:"limit" query:"limit"`
}

// GetAllMessages in channel
func GetAllMessages(c echo.Context) error {
	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	params := new(GetUpdatesRequestData)
	if err := c.Bind(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_params_error", "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	messages, err := chatService.GetMessages(uint(channelId), params.Limit)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}
