package app

import (
	"context"
	"github.com/deissh/rl/ayako/server"
	"github.com/deissh/rl/ayako/services"
	"github.com/deissh/rl/ayako/store"
	"time"
)

type App struct {
	srv     *server.Server
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

func runSecurityJob(s *App) {
	server.CreateRecurringTask("Security", func() {
		s.Srv().DoSecurityUpdateCheck()
	}, time.Hour*6)
}

func runUpdateCheck(s *App) {
	server.CreateRecurringTask("UpdateCheck", func() {
		s.Srv().DoBeatmapSetUpdate()
	}, time.Minute*30)
}

func runSearchNew(s *App) {
	server.CreateRecurringTask("SearchNew", func() {
		s.Srv().DoBeatmapSetSearchNew()
	}, time.Hour)
}

func (a *App) Context() context.Context {
	return a.context
}
func (a *App) Srv() *server.Server {
	return a.srv
}
func (a *App) Store() store.Store {
	return a.srv.Store
}

func (a *App) SetContext(c context.Context) {
	a.context = c
}
func (a *App) SetServer(srv *server.Server) {
	a.srv = srv
}
