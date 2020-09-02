package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"github.com/rs/zerolog/log"
)

// OAuthWithLog implements store.OAuth that is instrumented with zerolog
type OAuthWithLog struct {
	_base store.OAuth
}

func NewOAuthWithLog(base store.OAuth) store.OAuth {
	return OAuthWithLog{
		_base: base,
	}
}

// CreateClient implements store.OAuth
func (_d OAuthWithLog) CreateClient(ctx context.Context, name string, redirect string) (op1 *entity.OAuthClient, err error) {
	log.Trace().
		Interface("ctx", ctx).
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
	return _d._base.CreateClient(ctx, name, redirect)
}

// CreateToken implements store.OAuth
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

// GetClient implements store.OAuth
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

// RefreshToken implements store.OAuth
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

// RevokeToken implements store.OAuth
func (_d OAuthWithLog) RevokeToken(ctx context.Context, userId uint, accessToken string) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("accessToken", accessToken).
		Msg("store.OAuth.RevokeToken: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.RevokeToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.RevokeToken: finished")
		}
	}()
	return _d._base.RevokeToken(ctx, userId, accessToken)
}

// ValidateToken implements store.OAuth
func (_d OAuthWithLog) ValidateToken(ctx context.Context, accessToken string) (op1 *entity.OAuthToken, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("accessToken", accessToken).
		Msg("store.OAuth.ValidateToken: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.OAuth.ValidateToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.ValidateToken: finished")
		}
	}()
	return _d._base.ValidateToken(ctx, accessToken)
}
