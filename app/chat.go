package app

import (
	"context"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

var (
	ErrNotFoundChat = errors.New("chat", http.StatusNotFound, "Chat not found")
	ErrEmptyMessage = errors.New("chat", http.StatusBadRequest, "Message content must be not empty")
)

// CreateChat
func (a *App) CreateChat(ctx context.Context, userId, targetId uint, message string, isAction bool) (*entity.ChannelNewPm, error) {
	data, err := a.Store.Chat().CreatePM(ctx, userId, targetId, message, isAction)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// GetUpdatesInChat
func (a *App) GetUpdatesInChat(ctx context.Context, userId, since, channelId, limit uint) (*entity.ChannelUpdates, error) {
	if limit >= 100 || limit <= 0 {
		limit = 25
	}

	data, err := a.Store.Chat().GetUpdates(ctx, userId, since, channelId, limit)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// GetUpdatesInChat
func (a *App) GetMessages(ctx context.Context, userId, limit uint) (*[]entity.ChatMessage, error) {
	if limit >= 100 || limit <= 0 {
		limit = 25
	}

	data, err := a.Store.Chat().GetMessages(ctx, userId, limit)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// SendMessage to selected chat
func (a *App) SendMessage(ctx context.Context, userId, channelId uint, content string, isAction bool) (*entity.ChatMessage, error) {
	if content == "" {
		return nil, ErrEmptyMessage
	}

	data, err := a.Store.Chat().SendMessage(ctx, userId, channelId, content, isAction)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// GetAllPublicChats
func (a *App) GetAllPublicChats(ctx context.Context) (*[]entity.Channel, error) {
	data, err := a.Store.Chat().GetPublic(ctx)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// GetAllChats
func (a *App) GetAllChats(ctx context.Context, userId uint) (*[]entity.Channel, error) {
	data, err := a.Store.Chat().GetJoined(ctx, userId)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// JoinToChat
func (a *App) JoinToChat(ctx context.Context, userId, channelId uint) (*entity.Channel, error) {
	data, err := a.Store.Chat().Join(ctx, userId, channelId)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

// LeaveFromChat
func (a *App) LeaveFromChat(ctx context.Context, userId, channelId uint) error {
	err := a.Store.Chat().Leave(ctx, userId, channelId)
	if err != nil {
		return ErrNotFoundChat.WithCause(err)
	}

	return nil
}
