package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"github.com/rs/zerolog"
	"net/http"
)

type OAuthTokenController struct {
	UseCase *app.OAuthUseCase

	Logger *zerolog.Logger
}

var providerOAuthTokenSet = wire.NewSet(
	NewOAuthTokenController,
)

func NewOAuthTokenController(
	useCase *app.OAuthUseCase,
	logger *zerolog.Logger,
) *OAuthTokenController {
	return &OAuthTokenController{
		useCase,
		logger,
	}
}

func (h *OAuthTokenController) Create(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := &request.CreateOauthToken{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Token info not found")
	}

	token, err := h.UseCase.CreateOAuthToken(ctx, *params)
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
