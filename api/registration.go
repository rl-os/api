package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"net/http"
)

type RegistrationHandlers struct {
	App *app.App
}

func (h *RegistrationHandlers) Create(c echo.Context) error {
	params := request.CreateUser{}

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User info not found")
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user information")
	}

	ctx, _ := c.Get("context").(context.Context)

	user, err := h.App.Store.User().Create(ctx, params.Username, params.Email, params.Password)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
