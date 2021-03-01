package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"github.com/rs/zerolog"
	"strconv"
)

type ChatController struct {
	App    *app.App
	Logger *zerolog.Logger
}

var providerChatSet = wire.NewSet(
	NewChatController,
)

func NewChatController(app *app.App, logger *zerolog.Logger) *ChatController {
	return &ChatController{
		app,
		logger,
	}
}

// NewPm between 2 users
//
// @Router /api/v2/chat/new [post]
// @Tags Chat
// @Summary Create new PM channel between 2 users
// @Security OAuth2
// @Param payload body request.CreateNewChat true "JSON payload"
//
// @Success 200 {object} entity.ChannelNewPm
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) NewPm(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.CreateNewChat{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channels, err := h.App.CreateChat(ctx, userId, params.TargetId, params.Message, params.IsAction)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

// Updates in channel
//
// @Router /api/v2/chat/updates [get]
// @Tags Chat
// @Summary Returns updates in channel
// @Security OAuth2
// @Param since query integer true "since (last message id)"
// @Param channel_id query integer true "channel id"
// @Param limit query integer true "limit 1-100, default 50"
//
// @Success 200 {object} entity.ChannelUpdates
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) Updates(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.GetChatUpdates{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	updates, err := h.App.GetUpdatesInChat(
		ctx, userId, params.Since, params.ChannelId, params.Limit,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, updates)
}

// Messages in all joined chats
//
// @Router /api/v2/chat/channels/{id}/messages [get]
// @Tags Chat
// @Summary Returns Messages in all joined chats
// @Security OAuth2
// @Param limit query integer true "limit 1-100, default 50"
// @Param id path string false "channel id"
//
// @Success 200 {array} entity.ChatMessage
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) Messages(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.GetMessages{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	messages, err := h.App.GetMessages(
		ctx, userId, params.Limit,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

// Send message to chat
//
// @Router /api/v2/chat/channels/{id}/messages [post]
// @Tags Chat
// @Summary Send message to chat
// @Security OAuth2
// @Param payload body request.SendMessage true "JSON payload"
// @Param id path string false "channel id"
//
// @Success 200 {object} entity.ChatMessage
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) Send(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	params := request.SendMessage{}
	if err := c.Bind(params); err != nil {
		return errors.New("requires_params", 400, "invalid request params")
	}

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	messages, err := h.App.SendMessage(
		ctx, userId, uint(channelId), params.Message, params.IsAction,
	)
	if err != nil {
		return err
	}

	return c.JSON(200, messages)
}

// GetAll public chats
//
// @Router /api/v2/chat/channels [get]
// @Tags Chat
// @Summary Get all public chats
// @Security OAuth2
//
// @Success 200 {array} entity.Channel
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) GetAll(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	channels, err := h.App.GetAllPublicChats(ctx)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

// GetJoined chats
//
// @Router /api/v2/chat/channels/joined [get]
// @Tags Chat
// @Summary Get joined chats
// @Security OAuth2
//
// @Success 200 {array} entity.Channel
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) GetJoined(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channels, err := h.App.GetAllChats(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(200, channels)
}

// Join to chat
//
// @Router /api/v2/chat/channels/{id}/users/{user} [put]
// @Tags Chat
// @Summary Join to chat
// @Security OAuth2
// @Param id path string false "channel id"
// @Param user path string false "user id"
//
// @Success 200 {object} entity.Channel
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) Join(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	channel, err := h.App.JoinToChat(ctx, userId, uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, channel)
}

// Leave from chat
//
// @Router /api/v2/chat/channels/{id}/users/{user} [delete]
// @Tags Chat
// @Summary Leave from chat
// @Security OAuth2
// @Param id path string false "channel id"
// @Param user path string false "user id"
//
// @Success 200 {object} interface{}
// @Success 400 {object} errors.ResponseFormat
func (h *ChatController) Leave(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	channelId, err := strconv.ParseUint(c.Param("channel"), 10, 32)
	if err != nil {
		return errors.New("requires_params", 400, "invalid channelId")
	}

	err = h.App.LeaveFromChat(ctx, userId, uint(channelId))
	if err != nil {
		return err
	}

	return c.JSON(200, "{}")
}
