package client

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"sync"
	"time"
)

type OsuAPI struct {
	sync.Mutex

	domain       string
	clientId     string
	clientSecret string

	client       *gentleman.Client
	accessToken  string
	refreshToken string

	OAuth2     Oauth2API
	BeatmapSet BeatmapSetAPI
}

// WithAccessToken client
// auto refresh token if expire
func WithAccessToken(accessToken string, refreshToken string) OsuAPI {
	client := gentleman.New()

	api := OsuAPI{
		domain:       "https://osu.ppy.sh",
		clientId:     "5",
		clientSecret: "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk",

		client:       client,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}

	api.OAuth2 = Oauth2API{&api}
	api.BeatmapSet = BeatmapSetAPI{&api}

	client.BaseURL(api.domain)

	return api
}

// Bearer defines an authorization bearer token header in the outgoing request
func (client OsuAPI) bearerMiddleware() plugin.Plugin {
	return plugin.NewRequestPlugin(func(ctx *context.Context, h context.Handler) {
		accessToken := client.accessToken

		expiredAt, err := parseJwtExpiredAt(accessToken)
		if err != nil {
			h.Error(ctx, err)
			return
		}

		if time.Now().After(expiredAt) {
			token, err := client.OAuth2.TokenRenew("*", client.refreshToken)
			if token == nil || err != nil {
				h.Error(ctx, errors.Wrap(err, "failed token refresh"))
				return
			}

			client.accessToken = token.AccessToken
			client.refreshToken = token.RefreshToken
		}

		ctx.Request.Header.Set("Authorization", "Bearer "+client.accessToken)

		h.Next(ctx)
	})
}

// parseJwt and return claims
func parseJwtExpiredAt(token string) (time.Time, error) {
	parser := new(jwt.Parser)
	parsedToken, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return time.Now(), errors.Wrap(err, "invalid jwt format")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return time.Now(), errors.New("invalid claims format")
	}

	var tm time.Time
	switch iat := claims["iat"].(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}

	return tm, nil
}
