package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"github.com/rs/zerolog/log"
)

// OAuthWithLog implements repository.OAuth that is instrumented with zerolog
type OAuthWithLog struct {
	_base repository.OAuth
}

func NewOAuthWithLog(base repository.OAuth) repository.OAuth {
	return OAuthWithLog{
		_base: base,
	}
}

// CreateClient implements repository.OAuth
func (_d OAuthWithLog) CreateClient(ctx context.Context, userId uint, name string, redirect string) (op1 *entity.OAuthClient, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("name", name).
		Interface("redirect", redirect).
		Msg("store.OAuth.CreateClient: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.CreateClient: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.CreateClient: finished")
		}
	}()
	return _d._base.CreateClient(ctx, userId, name, redirect)
}

// CreateToken implements repository.OAuth
func (_d OAuthWithLog) CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (op1 *entity.OAuthToken, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("clientID", clientID).
		Interface("clientSecret", clientSecret).
		Interface("scopes", scopes).
		Msg("store.OAuth.CreateToken: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.CreateToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.CreateToken: finished")
		}
	}()
	return _d._base.CreateToken(ctx, userId, clientID, clientSecret, scopes)
}

// GetClient implements repository.OAuth
func (_d OAuthWithLog) GetClient(ctx context.Context, id uint, secret string) (op1 *entity.OAuthClient, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Interface("secret", secret).
		Msg("store.OAuth.GetClient: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.GetClient: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.GetClient: finished")
		}
	}()
	return _d._base.GetClient(ctx, id, secret)
}

// GetToken implements repository.OAuth
func (_d OAuthWithLog) GetToken(ctx context.Context, accessToken string) (op1 *entity.OAuthToken, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("accessToken", accessToken).
		Msg("store.OAuth.GetToken: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.GetToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.GetToken: finished")
		}
	}()
	return _d._base.GetToken(ctx, accessToken)
}

// RefreshToken implements repository.OAuth
func (_d OAuthWithLog) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (op1 *entity.OAuthToken, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("refreshToken", refreshToken).
		Interface("clientID", clientID).
		Interface("clientSecret", clientSecret).
		Msg("store.OAuth.RefreshToken: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.RefreshToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.RefreshToken: finished")
		}
	}()
	return _d._base.RefreshToken(ctx, refreshToken, clientID, clientSecret)
}

// RevokeAllTokens implements repository.OAuth
func (_d OAuthWithLog) RevokeAllTokens(ctx context.Context, userId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.OAuth.RevokeAllTokens: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.RevokeAllTokens: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.RevokeAllTokens: finished")
		}
	}()
	return _d._base.RevokeAllTokens(ctx, userId)
}
