package server

import (
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
)

// Server pass Store and RootRouter
type Server interface {
	GetStore() store.Store
	GetRootRouter() *echo.Group

	Start() error
	Shutdown() error
}
