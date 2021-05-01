package app

import (
	"context"
	"github.com/google/wire"
	"github.com/rl-os/api/services/cache"
	"github.com/rl-os/api/services/validator"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"time"
)

// ProviderAppSet provide DI
var ProviderAppSet = wire.NewSet(New, NewOptions)

// Options is app configuration struct
type Options struct {
	JWT struct {
		Secret       string
		BeforeRevoke time.Duration
	}
}

type App struct {
	Context context.Context
	Options *Options

	Validator *validator.Inst
	Cache     cache.Cache
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
	cache cache.Cache,
	options *Options,
	validator *validator.Inst,
) *App {
	ctx := context.TODO()

	return &App{
		Context:   ctx,
		Options:   options,
		Validator: validator,
		Cache:     cache,
	}
}

func (a *App) SetContext(ctx context.Context) {
	a.Context = ctx
}
