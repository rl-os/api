package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
)

type RegistrationHandlers struct {
	App *app.App
}

// Create new user
//
// @Router /users [post]
// @Summary Create new user
// @Description get string by ID
// @Param payload body request.CreateUser true "JSON payload"
//
// @Success 200 {object} entity.User
// @Success 400 {object} errors.ResponseFormat
func (h *RegistrationHandlers) Create(c echo.Context) error {
	params := &request.CreateUser{}

	if err := c.Bind(params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	user, err := h.App.Store.User().Create(ctx, params.Username, params.Email, params.Password)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
