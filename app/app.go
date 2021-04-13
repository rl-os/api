package app

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/rl-os/api/services"
	"github.com/rl-os/api/store"
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
	// Global context
	Context context.Context
	// Global validation context
	Validator *validator.Validate
	// Global configuration
	Options *Options

	// Store contains active implementation
	Store store.Store
	// Services that be enabled for this app
	Services *services.Services
}

// NewOptions create and parse config from viper instance
func NewOptions(v *viper.Viper) (*Options, error) {
	o := Options{}

	if err := v.UnmarshalKey("app", &o); err != nil {
		return nil, err
	}

	return &o, nil
}

// New with DI
func New(
	options *Options,
	store store.Store,
	services *services.Services,
) *App {
	app := &App{
		Store:     store,
		Options:   options,
		Services:  services,
		Validator: validator.New(),
	}

	return app
}

func (a *App) SetContext(ctx context.Context) {
	a.Context = ctx
}
