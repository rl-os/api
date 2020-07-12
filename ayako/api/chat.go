package api

import (
	"context"
	myctx "github.com/deissh/rl/ayako/ctx"
	"github.com/deissh/rl/ayako/errors"
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
	"strconv"
)

type ChatHandlers struct {
	Store store.Store
}

// newPmRequestData contain incoming data
type newPmRequestData struct {
	TargetId uint   `json:"target_id" query:"target_id" form:"target_id"`
	Message  string `json:"message" query:"message" form:"message"`
	IsAction bool   `json:"is_action" query:"is_action" form:"is_action"`
}

func (h *ChatHandlers) NewPm(c echo.Context) error {
	params := new(newPmRequestData)
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channels, err := h.Store.Chat().CreatePM(
		ctx, userId, params.TargetId, params.Message, params.IsAction,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

// getUpdatesRequestData contain incoming data
type getUpdatesRequestData struct {
	Since     uint `json:"since" query:"since"`
	ChannelId uint `json:"channel_id" query:"channel_id"`
	Limit     uint `json:"limit" query:"limit"`
}

func (h *ChatHandlers) Updates(c echo.Context) error {
	params := new(getUpdatesRequestData)
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	updates, err := h.Store.Chat().GetUpdates(
		ctx, userId, params.Since, params.ChannelId, params.Limit,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, updates)
}

// getMessagesRequestData contain incoming data
type getMessagesRequestData struct {
	Limit uint `json:"limit" query:"limit"`
}

func (h *ChatHandlers) Messages(c echo.Context) error {
	params := new(getMessagesRequestData)
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	messages, err := h.Store.Chat().GetMessages(
		ctx, userId, params.Limit,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

// sendMessageRequestData contain incoming data
type sendMessageRequestData struct {
	Message  string `json:"message" form:"message" validate:"required"`
	IsAction bool   `json:"is_action" form:"is_action"`
}

func (h *ChatHandlers) Send(c echo.Context) error {
	params := new(sendMessageRequestData)
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	messages, err := h.Store.Chat().SendMessage(
		ctx, userId, uint(channelId), params.Message, params.IsAction,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

func (h *ChatHandlers) GetAll(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	channels, err := h.Store.Chat().GetPublic(ctx)
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

	channels, err := h.Store.Chat().GetJoined(ctx, userId)
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

	channel, err := h.Store.Chat().Join(ctx, userId, uint(channelId))
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

	err = h.Store.Chat().Leave(ctx, userId, uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, "{}")
}
