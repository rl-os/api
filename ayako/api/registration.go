package api

import (
	"context"
	"github.com/deissh/rl/ayako/app"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegistrationHandlers struct {
	*app.App
}

func (h *RegistrationHandlers) Create(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	// CreateTokenRequestData contain incoming data with user credentials
	type CreateUserRequestData struct {
		Username string `json:"username" form:"user[username]" validate:"required"`
		Email    string `json:"email" form:"user[user_email]" validate:"required,email"`
		Password string `json:"password" form:"user[password]" validate:"required"`
	}

	params := new(CreateUserRequestData)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User info not found")
	}

	if err := validator.New().Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user information")
	}

	user, err := h.Store.User().Create(ctx, params.Username, params.Email, params.Password)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
