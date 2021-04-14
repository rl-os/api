package app

import (
	"context"
	"github.com/google/wire"
	"github.com/rl-os/api/services/bancho"
	"github.com/rl-os/api/store"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// ProviderSet provide DI
var ProviderSet = wire.NewSet(New, NewOptions)

// Options is app configuration struct
type Options struct {
	JWT struct {
		Secret string
	}
}

type App struct {
	Context context.Context
	Options *Options

	Store store.Store

	BanchoClient *bancho.Client
}

// NewOptions create and parse config from viper instance
func NewOptions(logger *zerolog.Logger, v *viper.Viper) (*Options, error) {
	o := Options{}

	logger.Debug().
		Msg("Loading config file")
	if err := v.UnmarshalKey("app", &o); err != nil {
		return nil, err
	}

	logger.Debug().
		Interface("app", o).
		Msg("Loaded config")

	return &o, nil
}

// New with DI
func New(
	options *Options,
	store store.Store,
	bancho *bancho.Client,
) *App {
	app := &App{
		Store:        store,
		Options:      options,
		BanchoClient: bancho,
	}

	return app
}

func (a *App) SetContext(ctx context.Context) {
	a.Context = ctx
}
