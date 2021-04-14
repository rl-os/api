package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/services/validator"
	"github.com/rs/zerolog"
	"net/http"
)

type FriendController struct {
	App       *app.App
	Logger    *zerolog.Logger
	Validator *validator.Inst
}

var providerFriendSet = wire.NewSet(
	NewFriendController,
)

func NewFriendController(
	app *app.App,
	logger *zerolog.Logger,
	validator *validator.Inst,
) *FriendController {
	return &FriendController{
		app,
		logger,
		validator,
	}
}

// GetAll friends/subscriptions
//
// @Router /api/v2/friends [get]
// @Tags Friends
// @Summary Return all friends/subscriptions
// @Security OAuth2
//
// @Success 200 {array} entity.UserShort
// @Success 400 {object} errors.ResponseFormat
func (h *FriendController) GetAll(c echo.Context) error {
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

// Add friend/subscription
//
// @Router /api/v2/friends [put]
// @Tags Friends
// @Summary Add friend/subscription
// @Security OAuth2
// @Param payload body request.FriendTargetId false "JSON payload"
// @Param target_id query string false "Target id (user)"
//
// @Success 200 {array} entity.UserShort
// @Success 400 {object} errors.ResponseFormat
func (h *FriendController) Add(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.FriendTargetId{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	if err := h.Validator.Struct(params); err != nil {
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

// Remove friend/subscription
//
// @Router /api/v2/friends [delete]
// @Tags Friends
// @Summary Remove friend/subscription
// @Security OAuth2
// @Param payload body request.FriendTargetId false "JSON payload"
// @Param target_id query string false "Target id (user)"
//
// @Success 200 {array} entity.UserShort
// @Success 400 {object} errors.ResponseFormat
func (h *FriendController) Remove(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.FriendTargetId{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	if err := h.Validator.Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate", err)
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := h.App.RemoveFriend(ctx, userId, params.TargetId)
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
