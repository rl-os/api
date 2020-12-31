package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"github.com/rs/zerolog/log"
)

var oauthDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "app_store_oauth_duration_seconds",
		Help:       "oauth runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

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
func (_d OAuthWithLog) CreateClient(ctx context.Context, userId uint, name string, redirect string) (op1 *entity.OAuthClient, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("name", name).
		Interface("redirect", redirect).
		Msg("store.OAuth.CreateClient: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.OAuth.CreateClient: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.CreateClient: finished")
		}
		oauthDurationSummaryVec.WithLabelValues("OAuth", "CreateClient", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.CreateClient(ctx, userId, name, redirect)
}

// CreateToken implements store.OAuth
func (_d OAuthWithLog) CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (op1 *entity.OAuthToken, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("clientID", clientID).
		Interface("clientSecret", clientSecret).
		Interface("scopes", scopes).
		Msg("store.OAuth.CreateToken: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.OAuth.CreateToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.CreateToken: finished")
		}
		oauthDurationSummaryVec.WithLabelValues("OAuth", "CreateToken", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.CreateToken(ctx, userId, clientID, clientSecret, scopes)
}

// GetClient implements store.OAuth
func (_d OAuthWithLog) GetClient(ctx context.Context, id uint, secret string) (op1 *entity.OAuthClient, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Interface("secret", secret).
		Msg("store.OAuth.GetClient: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.OAuth.GetClient: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.GetClient: finished")
		}
		oauthDurationSummaryVec.WithLabelValues("OAuth", "GetClient", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetClient(ctx, id, secret)
}

// GetToken implements store.OAuth
func (_d OAuthWithLog) GetToken(ctx context.Context, accessToken string) (op1 *entity.OAuthToken, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("accessToken", accessToken).
		Msg("store.OAuth.GetToken: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.OAuth.GetToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.GetToken: finished")
		}
		oauthDurationSummaryVec.WithLabelValues("OAuth", "GetToken", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetToken(ctx, accessToken)
}

// RefreshToken implements store.OAuth
func (_d OAuthWithLog) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (op1 *entity.OAuthToken, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("refreshToken", refreshToken).
		Interface("clientID", clientID).
		Interface("clientSecret", clientSecret).
		Msg("store.OAuth.RefreshToken: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.OAuth.RefreshToken: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.RefreshToken: finished")
		}
		oauthDurationSummaryVec.WithLabelValues("OAuth", "RefreshToken", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.RefreshToken(ctx, refreshToken, clientID, clientSecret)
}

// RevokeAllTokens implements store.OAuth
func (_d OAuthWithLog) RevokeAllTokens(ctx context.Context, userId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.OAuth.RevokeAllTokens: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.OAuth.RevokeAllTokens: returned an error")
		} else {
			log.Trace().
				Msg("store.OAuth.RevokeAllTokens: finished")
		}
		oauthDurationSummaryVec.WithLabelValues("OAuth", "RevokeAllTokens", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.RevokeAllTokens(ctx, userId)
}
