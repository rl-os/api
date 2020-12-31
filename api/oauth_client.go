package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity/request"
	"net/http"
)

type OAuthClientHandlers struct {
	App *app.App
}

// Create new oauth client
//
// @Router /api/v2/oauth/client [post]
// @Tags OAuth
// @Summary Create new oauth client
// @Security OAuth2
// @Param payload body request.CreateOAuthClient true "JSON payload"
//
// @Success 200 {object} entity.OAuthClient
// @Success 400 {object} errors.ResponseFormat
func (h OAuthClientHandlers) Create(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := &request.CreateOAuthClient{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Client info not found")
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid client information")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	client, err := h.App.CreateOAuthClient(ctx, userId, *params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, client)
}
