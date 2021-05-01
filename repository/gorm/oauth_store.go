package gorm

import (
	"context"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"gorm.io/gorm"
)

type OAuthRepository struct {
	*Supplier
}

func NewOAuthRepository(supplier *Supplier) repository.OAuth {
	return &OAuthRepository{supplier}
}

func (o OAuthRepository) CreateClient(ctx context.Context, userId uint, name string, redirect string) (*entity.OAuthClient, error) {
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

	return o.GetClient(ctx, client.ID, client.Secret)
}

func (o OAuthRepository) GetClient(ctx context.Context, id uint, secret string) (*entity.OAuthClient, error) {
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

func (o OAuthRepository) CreateToken(ctx context.Context, token entity.OAuthToken) (*entity.OAuthToken, error) {
	err := o.GetMaster().WithContext(ctx).
		Table("oauth_token").
		Create(&token).
		Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (o OAuthRepository) RevokeAllTokens(ctx context.Context, userId uint) error {
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

func (o OAuthRepository) RefreshToken(ctx context.Context, refreshToken string) (*entity.OAuthToken, error) {
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

	return o.CreateToken(ctx, token)
}

func (o OAuthRepository) GetToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
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
