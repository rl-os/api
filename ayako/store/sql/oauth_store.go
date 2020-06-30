package sql

import (
	"context"
	"github.com/deissh/go-utils"
	myctx "github.com/deissh/rl/ayako/ctx"
	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/errors"
	"github.com/deissh/rl/ayako/store"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type OAuthStore struct {
	SqlStore
}

func newSqlOAuthStore(sqlStore SqlStore) store.OAuth {
	return &OAuthStore{sqlStore}
}

func (o OAuthStore) CreateClient(ctx context.Context, name string, redirect string) (*entity.OAuthClient, error) {
	var client entity.OAuthClient

	secret, err := utils.GenerateRandomString(255)
	if err != nil {
		return nil, errors.WithCause(500, "Internal error", err)
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return nil, errors.WithCause(401, "Require user_id", err)
	}

	err = o.GetMaster().GetContext(
		ctx,
		&client,
		`INSERT INTO oauth_client (user_id, name, secret, redirect)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
		userId, name, secret, redirect,
	)

	return &client, nil
}

func (o OAuthStore) GetClient(ctx context.Context, id uint, secret string) (*entity.OAuthClient, error) {
	client := entity.OAuthClient{}

	err := o.GetMaster().GetContext(
		ctx,
		&client,
		`SELECT * FROM oauth_client WHERE id = $1 AND secret = $2`,
		id, secret,
	)

	return &client, err
}

func (o OAuthStore) CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (*entity.OAuthToken, error) {
	_, err := o.OAuth().GetClient(ctx, clientID, clientSecret)
	if err != nil {
		return nil, errors.WithCause(404, "OAuth client not found", err)
	}

	cfg := o.GetConfig()

	token, err := tokenCreate(
		cfg.JWT.Secret,
		time.Duration(int(time.Hour)*cfg.JWT.HoursBeforeRevoke),
		userId,
		clientID,
		scopes,
	)
	if err != nil {
		return nil, err
	}

	err = o.GetMaster().GetContext(
		ctx,
		token,
		`INSERT INTO oauth_token (user_id, client_id, access_token, refresh_token, scopes, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		token.UserID, token.ClientID, token.AccessToken,
		token.RefreshToken, token.Scopes, token.ExpiresAt,
	)
	if err != nil {
		return nil, errors.WithCause(500, "creating token", err)
	}

	return token, nil
}

func (o OAuthStore) RevokeToken(ctx context.Context, userId uint, accessToken string) error {
	_, err := o.GetMaster().ExecContext(
		ctx,
		`UPDATE oauth_token
			SET revoked = true
			WHERE user_id = $1 AND client_id = $2 AND revoked = false
			RETURNING *`,
		userId,
		accessToken,
	)
	if err != nil {
		return errors.WithCause(404, "Access token not exist or already revoked", err)
	}

	return nil
}

func (o OAuthStore) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error) {
	var token entity.OAuthToken
	err := o.GetMaster().GetContext(
		ctx,
		&token,
		`UPDATE oauth_token
			SET revoked = true
			WHERE refresh_token = $1 AND revoked = false AND client_id = $2
			RETURNING *`,
		refreshToken,
		clientID,
	)
	if err != nil {
		return nil, errors.WithCause(404, "Access token not exist or already revoked", err)
	}

	newToken, err := o.OAuth().CreateToken(ctx, token.UserID, clientID, clientSecret, token.Scopes)
	return newToken, err
}

func (o OAuthStore) ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
	cfg := o.GetConfig()

	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	v, _ := err.(*jwt.ValidationError)
	if err != nil && v.Errors == jwt.ValidationErrorExpired {
		return nil, errors.WithCause(400, "Access token expired", err)
	} else if err != nil {
		return nil, errors.WithCause(401, "Invalid access token", err)
	}

	var token entity.OAuthToken
	err = o.GetMaster().GetContext(
		ctx,
		&token,
		`SELECT * FROM oauth_token
				WHERE access_token = $1`,
		accessToken,
	)
	if err != nil {
		return nil, errors.WithCause(500, "Selecting access_token in database errors.", err)
	}
	if token.Revoked {
		return nil, errors.WithCause(400, "Access token expired", err)
	}

	return &token, nil
}
