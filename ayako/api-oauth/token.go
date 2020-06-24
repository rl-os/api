package api_oauth

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
)

type TokenHandlers struct {
	Store store.Store
}

func (api *Routes) InitToken(store store.Store) {
	handlers := TokenHandlers{store}

	api.Token.POST("/", handlers.Create)
}

func (h *TokenHandlers) Create(c echo.Context) error {
	return c.JSON(200, nil)
}
