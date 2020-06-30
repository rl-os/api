package chats

import (
	"github.com/deissh/rl/api/pkg"
	"github.com/deissh/rl/api/pkg/entity"
	chatService "github.com/deissh/rl/api/pkg/services/chat"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// SendMessageRequestData contain incoming data
type SendMessageRequestData struct {
	Message  string `json:"message" form:"message" validate:"required"`
	IsAction bool   `json:"is_action" form:"is_action"`
}

// ChannelSendHandler create new message
func ChannelSendHandler(c echo.Context) (err error) {
	data := c.Get("current_user")

	current, ok := data.(*entity.User)
	if !ok {
		return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
	}

	params := new(SendMessageRequestData)
	if err := c.Bind(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_params_error", "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err := chatService.SendMessage(current.ID, uint(channelId), params.Message, params.IsAction)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
