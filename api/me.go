package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
)

type MeHandlers struct {
	App *app.App
}

// Me, current user
//
// @Router /api/v2/me/{mode} [get]
// @Tags Current user
// @Summary Return current user
// @Security OAuth2
// @Param mode path string false "game mod"
//
// @Success 200 {object} entity.User
// @Success 400 {object} errors.ResponseFormat
func (h *MeHandlers) Me(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)
	mode := c.Param("mode")

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	user, err := h.App.User.Get(ctx, userId, mode)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
