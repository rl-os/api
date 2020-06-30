package chats

import (
	"github.com/deissh/rl/api/pkg"
	"github.com/deissh/rl/api/pkg/entity"
	chatService "github.com/deissh/rl/api/pkg/services/chat"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetUpdatesRequestData contain incoming data
type GetUpdatesRequestData struct {
	Since     uint `json:"since" query:"since"`
	ChannelId uint `json:"channel_id" query:"channel_id"`
	Limit     uint `json:"limit" query:"limit"`
}

// ChannelUpdatesHandler return updates in channel
func ChannelUpdatesHandler(c echo.Context) (err error) {
	data := c.Get("current_user")

	current, ok := data.(*entity.User)
	if !ok {
		return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
	}

	params := new(GetUpdatesRequestData)
	if err := c.Bind(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_params_error", "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	updates, err := chatService.GetUpdates(current.ID, params.Since, params.ChannelId, params.Limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updates)
}
