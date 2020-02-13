package chats

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	chatService "github.com/deissh/osu-lazer/api/pkg/services/chat"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetUpdatesRequestData contain incoming data
type NewPmRequestData struct {
	TargetId uint   `json:"target_id" query:"target_id" form:"target_id"`
	Message  string `json:"message" query:"message" form:"message"`
	IsAction bool   `json:"is_action" query:"is_action" form:"is_action"`
}

// NewPm channel
func NewPm(c echo.Context) error {
	data := c.Get("current_user")

	current, ok := data.(*entity.User)
	if !ok {
		return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
	}

	params := new(NewPmRequestData)
	if err := c.Bind(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_params_error", "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	newPm, err := chatService.NewPm(current.ID, params.TargetId, params.Message, params.IsAction)
	if err != nil {
		return err
	}

	return c.JSON(200, newPm)
}
