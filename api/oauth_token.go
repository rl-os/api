package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"net/http"
)

type OAuthTokenHandlers struct {
	App *app.App
}

func (h *OAuthTokenHandlers) Create(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.CreateOauthToken{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Token info not found")
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token information")
	}

	var token *entity.OAuthToken

	if params.GrantType == "password" {
		user, err := h.App.Store.User().GetByBasic(ctx, params.Username, params.Password)
		if err != nil {
			return err
		}

		token, err = h.App.Store.OAuth().CreateToken(
			ctx,
			user.ID,
			params.ClientID,
			params.ClientSecret,
			params.Scope,
		)
		if err != nil {
			return err
		}
	} else if params.GrantType == "refresh_token" {
		var err error
		token, err = h.App.Store.OAuth().RefreshToken(ctx, params.RefreshToken, params.ClientID, params.ClientSecret)
		if err != nil {
			return err
		}
	} else {
		return errors.New("oauth_token_create", 400, "Invalid grand_type")
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
