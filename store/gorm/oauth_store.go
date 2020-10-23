package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
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
	panic("implement me")
}

func (o OAuthStore) CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (*entity.OAuthToken, error) {
	panic("implement me")
}

func (o OAuthStore) RevokeToken(ctx context.Context, userId uint, accessToken string) error {
	panic("implement me")
}

func (o OAuthStore) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error) {
	panic("implement me")
}

func (o OAuthStore) ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
	panic("implement me")
}
