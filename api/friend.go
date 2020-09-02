package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"net/http"
)

type FriendHandlers struct {
	App *app.App
}

func (h *FriendHandlers) GetAll(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := h.App.Store.Friend().GetSubscriptions(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(200, users)
}

func (h *FriendHandlers) Add(c echo.Context) error {
	// contain incoming data
	type putFriendRequestData struct {
		TargetId uint `json:"target_id" query:"target_id"`
	}

	params := new(putFriendRequestData)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.App.Store.Friend().Add(ctx, userId, params.TargetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	users, err := h.App.Store.Friend().GetSubscriptions(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(200, users)
}

func (h *FriendHandlers) Remove(c echo.Context) error {
	type requestData struct {
		TargetId uint `json:"target_id" query:"target_id"`
	}

	params := new(requestData)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.App.Store.Friend().Remove(ctx, userId, params.TargetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	users, err := h.App.Store.Friend().GetSubscriptions(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(200, users)
}
