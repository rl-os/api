package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"net/http"
)

type OAuthTokenHandlers struct {
	App *app.App
}

func (h *OAuthTokenHandlers) Create(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := &request.CreateOauthToken{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Token info not found")
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token information")
	}

	token, err := h.App.OAuth.CreateToken(ctx, *params)
	if err != nil {
		return err
	}

	return c.JSON(200, struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
	}{
		token.AccessToken,
		token.RefreshToken,
		"Bearer",
		token.ExpiresIn,
	})
}
