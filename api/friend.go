package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity/request"
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

	users, err := h.App.GetAllFriends(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}

func (h *FriendHandlers) Add(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.FriendTargetId{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := h.App.AddFriend(ctx, userId, params.TargetId)
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}

func (h *FriendHandlers) Remove(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.FriendTargetId{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	if err := h.App.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := h.App.RemoveFriend(ctx, userId, params.TargetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(200, users)
}
