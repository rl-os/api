package app

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/services"
	"github.com/rl-os/api/store"
)

type App struct {
	// Global context
	Context context.Context
	// Global validation context
	Validator *validator.Validate
	// Global configuration
	Config *config.Config

	// Store contains active implementation
	Store store.Store
	// Services that be enabled for this app
	Services *services.Services
}

// NewApp with DI
func NewApp(
	store store.Store,
	config *config.Config,
	services *services.Services,
) *App {
	app := &App{
		Store:     store,
		Config:    config,
		Services:  services,
		Validator: validator.New(),
	}

	return app
}

func (a *App) SetContext(ctx context.Context) {
	a.Context = ctx
}
