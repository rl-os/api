package sql

import (
	"context"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"gorm.io/gorm"
	"time"
)

type OAuthStore struct {
	SqlStore
}

func newSqlOAuthStore(sqlStore SqlStore) store.OAuth {
	return &OAuthStore{sqlStore}
}

func (o OAuthStore) CreateClient(ctx context.Context, userId uint, name string, redirect string) (*entity.OAuthClient, error) {
	secret, err := utils.GenerateRandomString(64)
	if err != nil {
		return nil, err
	}

	client := entity.OAuthClient{
		UserID:   userId,
		Name:     name,
		Secret:   secret,
		Redirect: redirect,
		Revoked:  false,
	}

	err = o.GetMaster().Transaction(func(tx *gorm.DB) error {
		return tx.WithContext(ctx).
			Table("oauth_client").
			Create(&client).
			Error
	})

	if err != nil {
		return nil, err
	}

	return o.OAuth().GetClient(ctx, client.ID, client.Secret)
}

func (o OAuthStore) GetClient(ctx context.Context, id uint, secret string) (*entity.OAuthClient, error) {
	client := entity.OAuthClient{}

	err := o.GetMaster().
		WithContext(ctx).
		Table("oauth_client").
		Where("id = ? AND secret = ?", id, secret).
		First(&client).
		Error
	if err != nil {
		return nil, err
	}

	return &client, err
}

func (o OAuthStore) CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (*entity.OAuthToken, error) {
	_, err := o.OAuth().GetClient(ctx, clientID, clientSecret)
	if err != nil {
		return nil, err
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

	err = o.GetMaster().WithContext(ctx).
		Table("oauth_token").
		Create(&token).
		Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (o OAuthStore) RevokeAllTokens(ctx context.Context, userId uint) error {
	err := o.GetMaster().
		WithContext(ctx).
		Table("oauth_token").
		Where("user_id = ?", userId).
		Update("revoked", true).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (o OAuthStore) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error) {
	token := entity.OAuthToken{}

	err := o.GetMaster().
		WithContext(ctx).
		Table("oauth_token").
		Where("refresh_token = ?", refreshToken).
		Update("revoked", true).
		First(&token).
		Error
	if err != nil {
		return nil, err
	}

	return o.OAuth().CreateToken(ctx, token.UserID, token.ClientID, clientSecret, token.Scopes)
}

func (o OAuthStore) GetToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
	token := entity.OAuthToken{}

	err := o.GetMaster().
		WithContext(ctx).
		Table("oauth_token").
		Where("access_token = ?", accessToken).
		First(&token).
		Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}
