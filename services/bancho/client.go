package bancho

import (
	"errors"
	osu "github.com/deissh/osu-go-client"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(NewOptions, New)

var (
	AuthErr = errors.New("invalid login and password or oauth client")
)

type Options struct {
	Username string
	Password string

	OAuthClient struct {
		Id     string
		Secret string
	} `mapstructure:"oauth_client"`
}

type Client struct {
	*osu.OsuAPI
}

// NewOptions create and parse config from viper instance
func NewOptions(logger *zerolog.Logger, v *viper.Viper) (*Options, error) {
	o := Options{}

	logger.Debug().
		Msg("Loading config file")
	if err := v.UnmarshalKey("bancho", &o); err != nil {
		return nil, err
	}

	logger.Debug().
		Interface("bancho", o).
		Msg("Loaded config")

	return &o, nil
}

func New(options *Options) *Client {
	if options.OAuthClient.Id != "" || options.OAuthClient.Secret != "" {
		osu.APIClientId = options.OAuthClient.Id
		osu.APIClientSecret = options.OAuthClient.Secret
	}

	client, err := osu.WithBasicAuth(
		options.Username,
		options.Password,
	)

	if err != nil {
		log.Warn().
			Err(AuthErr).
			Err(err).
			Send()

		return nil
	}

	return &Client{client}
}
