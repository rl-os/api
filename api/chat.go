package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"strconv"
)

type ChatHandlers struct {
	App *app.App
}

func (h *ChatHandlers) NewPm(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.CreateNewChat{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channels, err := h.App.Store.Chat().CreatePM(
		ctx, userId, params.TargetId, params.Message, params.IsAction,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

func (h *ChatHandlers) Updates(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.GetChatUpdates{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	updates, err := h.App.Store.Chat().GetUpdates(
		ctx, userId, params.Since, params.ChannelId, params.Limit,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, updates)
}

func (h *ChatHandlers) Messages(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.GetMessages{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	messages, err := h.App.Store.Chat().GetMessages(
		ctx, userId, params.Limit,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

func (h *ChatHandlers) Send(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.SendMessage{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	messages, err := h.App.Store.Chat().SendMessage(
		ctx, userId, uint(channelId), params.Message, params.IsAction,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

func (h *ChatHandlers) GetAll(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	channels, err := h.App.Store.Chat().GetPublic(ctx)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

func (h *ChatHandlers) GetJoined(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channels, err := h.App.Store.Chat().GetJoined(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

func (h *ChatHandlers) Join(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	// todo: parse from request and if current user is admin remove by userId
	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	channel, err := h.App.Store.Chat().Join(ctx, userId, uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, channel)
}

func (h *ChatHandlers) Leave(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	err = h.App.Store.Chat().Leave(ctx, userId, uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, "{}")
}
