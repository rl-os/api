package app

import (
	"context"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity/request"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

var (
	OAuthGroundType      = []string{"password", "refresh_token"}
	InvalidAuthParamsErr = errors.New("invalid_auth_params", http.StatusBadRequest, "Not found")
)

// CreateOAuthToken and return it
func (a *App) CreateOAuthToken(ctx context.Context, request request.CreateOauthToken) (*entity.OAuthToken, error) {
	if !utils.ContainsString(OAuthGroundType, request.GrantType) {
		request.GrantType = "password"
	}

	if request.GrantType == "password" {
		user, err := a.Store.User().GetByBasic(
			ctx,
			request.Username,
			request.Password,
		)
		if err != nil {
			return nil, InvalidAuthParamsErr.WithCause(err)
		}

		token, err := a.Store.OAuth().CreateToken(
			ctx,
			user.ID,
			request.ClientID,
			request.ClientSecret,
			request.Scope,
		)
		if err != nil {
			return nil, InvalidAuthParamsErr.WithCause(err)
		}

		return token, nil
	} else if request.GrantType == "refresh_token" {
		var err error
		token, err := a.Store.OAuth().RefreshToken(ctx, request.RefreshToken, request.ClientID, request.ClientSecret)
		if err != nil {
			return nil, InvalidAuthParamsErr.WithCause(err)
		}

		return token, nil
	}

	return nil, InvalidAuthParamsErr.WithCause(nil, "invalid grant type")
}
