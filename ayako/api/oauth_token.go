package api

import (
	"context"
	"github.com/deissh/rl/ayako/app"
	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OAuthTokenHandlers struct {
	*app.App
}

func (h *OAuthTokenHandlers) Create(c echo.Context) error {
	// createTokenRequestData contain incoming data with user credentials
	type createTokenRequestData struct {
		GrantType    string `json:"grant_type" form:"grant_type" validate:"required"`
		ClientID     uint   `json:"client_id" form:"client_id" validate:"required"`
		ClientSecret string `json:"client_secret" form:"client_secret" validate:"required"`
		Scope        string `json:"scope" form:"scope" validate:"required"`

		// grant_type == password
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`

		// grant_type == refresh_token
		RefreshToken string `json:"refresh_token" form:"refresh_token"`
	}

	ctx, _ := c.Get("context").(context.Context)

	params := new(createTokenRequestData)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Token info not found")
	}

	if err := validator.New().Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid token information")
	}

	var token *entity.OAuthToken

	if params.GrantType == "password" {
		user, err := h.Store.User().GetByBasic(ctx, params.Username, params.Password)
		if err != nil {
			return err
		}

		token, err = h.Store.OAuth().CreateToken(
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
		token, err = h.Store.OAuth().RefreshToken(ctx, params.RefreshToken, params.ClientID, params.ClientSecret)
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
