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
	// содержит доступные способы авторизации через oauth2
	// password - авторизация с помощью логина и
	// пароля которые передаются в теле запроса
	// refresh_token - авторизация с использованием
	// рефреш токена который был получен во время предыдущей авторизации
	OAuthGroundType = []string{"password", "refresh_token"}

	// Список ошибок которые может вернуть данный модуль
	// необходим для того что бы не явно не повторять
	// уникальный индификатор ошибки
	InvalidAuthParamsErr        = errors.New("invalid_auth_params", http.StatusBadRequest, "Not found")
	InvalidAuthTokenErr         = errors.New("oauth_validate_token", http.StatusUnauthorized, "Invalid access token")
	ExpiredAuthTokenErr         = errors.New("oauth_expire_token", http.StatusUnauthorized, "Access token expired")
	NotFoundRefreshAuthErr      = errors.New("oauth_invalid_refresh", http.StatusBadRequest, "Not found")
	InvalidOAuthClientParamsErr = errors.New("oauth_invalid_client_params", http.StatusBadRequest, "Invalid oauth client")
)

type OAuth struct {
	*App
}

// ValidateToken проверяет переданный функции токен доступа,
// если все нормально и токен еще не истек то вернется entity.OAuthToken
// при ошибке может вернуть InvalidAuthTokenErr или ExpiredAuthTokenErr
func (a *OAuth) ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
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

// CreateClient создает новый entity.OAuthClient
// в случае InvalidOAuthClientParamsErr
func (a OAuth) CreateClient(ctx context.Context, userId uint, request request.CreateOAuthClient) (*entity.OAuthClient, error) {
	token, err := a.Store.OAuth().CreateClient(ctx, userId, request.Name, request.Redirect)
	if err != nil {
		return nil, InvalidOAuthClientParamsErr.WithCause(err)
	}

	return token, nil
}

// RefreshToken and revoke old access token
func (a *OAuth) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error) {
	token, err := a.Store.OAuth().RefreshToken(ctx, refreshToken, clientID, clientSecret)
	if err != nil {
		return nil, NotFoundRefreshAuthErr.WithCause(err)
	}

	return token, nil
}

// CreateToken and return it
func (a *OAuth) CreateToken(ctx context.Context, request request.CreateOauthToken) (*entity.OAuthToken, error) {
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
		return a.RefreshToken(
			ctx,
			request.RefreshToken,
			request.ClientID,
			request.ClientSecret,
		)
	}

	return nil, InvalidAuthParamsErr.WithCause(nil, "invalid grant type")
}
