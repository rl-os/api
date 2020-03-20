package client

import (
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"time"
)

const (
	Domain       = "https://osu.ppy.sh"
	ClientId     = "5"
	ClientSecret = "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk"
)

type OsuAPI struct {
	Client       *gentleman.Client
	AccessToken  string
	RefreshToken string
	RefreshRate  time.Duration

	BeatmapSet BeatmapSetAPI
}

func WithAccessToken(accessToken string, refreshToken string) OsuAPI {
	client := gentleman.New()
	client.BaseURL(Domain)

	api := OsuAPI{
		Client:       client,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	client.Use(Bearer(api))

	api.BeatmapSet = BeatmapSetAPI{&api}

	return api
}

// Bearer defines an authorization bearer token header in the outgoing request
func Bearer(client OsuAPI) plugin.Plugin {
	return plugin.NewRequestPlugin(func(ctx *context.Context, h context.Handler) {
		ctx.Request.Header.Set("Authorization", "Bearer "+client.AccessToken)
		h.Next(ctx)
	})
}
