package api_oauth

import (
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Token        *echo.Group
	Registration *echo.Group
}

func New(store store.Store, prefix *echo.Group) {
	api := Routes{}

	api.Token = prefix.Group("/oauth/token")
	api.Registration = prefix.Group("/users")

	api.InitToken(store)
	api.InitRegistration(store)
}

func empty(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{
		"message": "I'm alive!",
	})
}
