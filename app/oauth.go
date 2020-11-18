package app

import (
	"context"
	"github.com/deissh/go-utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/rl-os/api/entity/request"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

var (
	OAuthGroundType      = []string{"password", "refresh_token"}
	InvalidAuthParamsErr = errors.New("invalid_auth_params", http.StatusBadRequest, "Not found")
	InvalidAuthTokenErr  = errors.New("oauth_validate_token", http.StatusUnauthorized, "Invalid access token")
	ExpiredAuthTokenErr  = errors.New("oauth_expire_token", http.StatusUnauthorized, "Access token expired")
)

func (a *App) ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.Config.JWT.Secret), nil
	})

	if err != nil {
		return nil, InvalidAuthTokenErr.WithCause(err)
	}

	token, err := a.Store.OAuth().GetToken(ctx, accessToken)
	if err != nil {
		return nil, InvalidAuthTokenErr.WithCause(err, "Not found")
	}

	if token.Revoked {
		return nil, ExpiredAuthTokenErr
	}

	return token, nil
}

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
