package app

import (
	"context"
	"github.com/deissh/rl/ayako/server"
	"github.com/deissh/rl/ayako/services"
	"github.com/deissh/rl/ayako/store"
)

type App struct {
	srv     server.Server
	context context.Context

	services *services.Services
}

// NewApp with DI
func NewApp(services *services.Services) *App {
	app := &App{
		services: services,
	}

	return app
}

func (a *App) Context() context.Context {
	return a.context
}
func (a *App) Srv() server.Server {
	return a.srv
}
func (a *App) Store() store.Store {
	return a.srv.GetStore()
}

func (a *App) SetContext(c context.Context) {
	a.context = c
}
func (a *App) SetServer(srv server.Server) {
	a.srv = srv
}
