package api

import (
	"context"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MeHandlers struct {
	Store store.Store
}

func (api *Routes) InitMe(store store.Store) {
	handlers := MeHandlers{store}

	api.Me.GET("/", handlers.Me)
}

func (h *MeHandlers) Me(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	//userId, err	 := myctx.GetUserID(ctx)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	//}

	mode := c.Param("mode")

	user, err := h.Store.User().Get(ctx, 1, mode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(200, user)
}
