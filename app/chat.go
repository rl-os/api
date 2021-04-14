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

	InvalidPMBodyErr      = errors.New("create_pm", http.StatusBadRequest, "Invalid chat information")
	InvalidMessageBodyErr = errors.New("send_message_chat", http.StatusBadRequest, "Invalud message body")
)

func (a *App) CreateChat(ctx context.Context, userId, targetId uint, message string, isAction bool) (*entity.ChannelNewPm, error) {
	channel, err := a.ChatRepository.CreatePm(ctx, userId, targetId)
	if err != nil {
		return nil, InvalidPMBodyErr.WithCause(err)
	}

	msg, err := a.ChatRepository.SendMessage(ctx, userId, channel.ID, message, isAction)
	if err != nil {
		return nil, InvalidMessageBodyErr.WithCause(err)
	}

	presence, err := a.ChatRepository.GetJoined(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &entity.ChannelNewPm{
		Id:       channel.ID,
		Presence: *presence,
		Messages: *msg,
	}, nil
}

func (a *App) GetUpdatesInChat(ctx context.Context, userId, since, channelId, limit uint) (*entity.ChannelUpdates, error) {
	if limit >= 100 || limit <= 0 {
		limit = 25
	}

	channels, err := a.ChatRepository.GetJoined(ctx, userId)
	if err != nil {
		return nil, err
	}

	messages, err := a.ChatRepository.GetMessages(ctx, userId, since, limit)
	if err != nil {
		return nil, err
	}

	return &entity.ChannelUpdates{
		Presence: channels,
		Messages: messages,
	}, nil
}

// GetUpdatesInChat
func (a *App) GetMessages(ctx context.Context, userId, limit uint) (*[]entity.ChatMessage, error) {
	if limit >= 100 || limit <= 0 {
		limit = 25
	}

	data, err := a.ChatRepository.GetMessages(ctx, userId, limit, limit)
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

	data, err := a.ChatRepository.SendMessage(ctx, userId, channelId, content, isAction)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

func (a *App) GetAllPublicChats(ctx context.Context) (*[]entity.Channel, error) {
	data, err := a.ChatRepository.GetPublic(ctx)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

func (a *App) GetAllChats(ctx context.Context, userId uint) (*[]entity.Channel, error) {
	data, err := a.ChatRepository.GetJoined(ctx, userId)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

func (a *App) JoinToChat(ctx context.Context, userId, channelId uint) (*entity.Channel, error) {
	data, err := a.ChatRepository.Join(ctx, userId, channelId)
	if err != nil {
		return nil, ErrNotFoundChat.WithCause(err)
	}

	return data, nil
}

func (a *App) LeaveFromChat(ctx context.Context, userId, channelId uint) error {
	err := a.ChatRepository.Leave(ctx, userId, channelId)
	if err != nil {
		return ErrNotFoundChat.WithCause(err)
	}

	return nil
}
