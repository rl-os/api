package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"time"
)

type OAuthStore struct {
	SqlStore
}

func newSqlOAuthStore(sqlStore SqlStore) store.OAuth {
	return &OAuthStore{sqlStore}
}

func (o OAuthStore) CreateClient(ctx context.Context, name string, redirect string) (*entity.OAuthClient, error) {
	panic("implement me")
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

func (o OAuthStore) RevokeToken(ctx context.Context, userId uint, accessToken string) error {
	panic("implement me")
}

func (o OAuthStore) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error) {
	panic("implement me")
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
