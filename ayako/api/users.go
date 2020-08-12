package api

import (
	"context"
	"github.com/deissh/rl/ayako/app"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UsersHandlers struct {
	*app.App
}

func (h *UsersHandlers) Get(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mode := c.Param("mode")

	user, err := h.Store.User().Get(ctx, uint(userId), mode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(200, user)
}
