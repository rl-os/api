package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/deissh/osu-lazer/ayako/store"
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
func (_d OAuthWithLog) CreateClient(ctx context.Context, name string, redirect string) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("name", name).
		Interface("redirect", redirect).
		Msg("store.OAuth.CreateClient: calling")
	defer func() {
		log.Debug().
			Msg("store.OAuth.CreateClient: finished")
	}()
	_d._base.CreateClient(ctx, name, redirect)
	return
}

// CreateToken implements store.OAuth
func (_d OAuthWithLog) CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("clientID", clientID).
		Interface("clientSecret", clientSecret).
		Interface("scopes", scopes).
		Msg("store.OAuth.CreateToken: calling")
	defer func() {
		log.Debug().
			Msg("store.OAuth.CreateToken: finished")
	}()
	_d._base.CreateToken(ctx, userId, clientID, clientSecret, scopes)
	return
}

// GetClient implements store.OAuth
func (_d OAuthWithLog) GetClient(ctx context.Context, id uint) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.OAuth.GetClient: calling")
	defer func() {
		log.Debug().
			Msg("store.OAuth.GetClient: finished")
	}()
	_d._base.GetClient(ctx, id)
	return
}

// RefreshToken implements store.OAuth
func (_d OAuthWithLog) RefreshToken(ctx context.Context, refreshToken string) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("refreshToken", refreshToken).
		Msg("store.OAuth.RefreshToken: calling")
	defer func() {
		log.Debug().
			Msg("store.OAuth.RefreshToken: finished")
	}()
	_d._base.RefreshToken(ctx, refreshToken)
	return
}

// RevokeToken implements store.OAuth
func (_d OAuthWithLog) RevokeToken(ctx context.Context, userId uint, accessToken string) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("accessToken", accessToken).
		Msg("store.OAuth.RevokeToken: calling")
	defer func() {
		log.Debug().
			Msg("store.OAuth.RevokeToken: finished")
	}()
	_d._base.RevokeToken(ctx, userId, accessToken)
	return
}

// ValidateToken implements store.OAuth
func (_d OAuthWithLog) ValidateToken(ctx context.Context, accessToken string) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("accessToken", accessToken).
		Msg("store.OAuth.ValidateToken: calling")
	defer func() {
		log.Debug().
			Msg("store.OAuth.ValidateToken: finished")
	}()
	_d._base.ValidateToken(ctx, accessToken)
	return
}
