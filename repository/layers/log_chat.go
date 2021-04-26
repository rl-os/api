package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"github.com/rs/zerolog/log"
)

// ChatWithLog implements repository.Chat that is instrumented with zerolog
type ChatWithLog struct {
	_base repository.Chat
}

func NewChatWithLog(base repository.Chat) repository.Chat {
	return ChatWithLog{
		_base: base,
	}
}

// CreatePm implements repository.Chat
func (_d ChatWithLog) CreatePm(ctx context.Context, userId uint, targetId uint) (cp1 *entity.Channel, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Chat.CreatePm: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.CreatePm: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.CreatePm: finished")
		}
	}()
	return _d._base.CreatePm(ctx, userId, targetId)
}

// Get implements repository.Chat
func (_d ChatWithLog) Get(ctx context.Context, channelId uint) (cp1 *entity.Channel, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("channelId", channelId).
		Msg("store.Chat.Get: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.Get: finished")
		}
	}()
	return _d._base.Get(ctx, channelId)
}

// GetJoined implements repository.Chat
func (_d ChatWithLog) GetJoined(ctx context.Context, userId uint) (cap1 *[]entity.Channel, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Chat.GetJoined: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.GetJoined: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetJoined: finished")
		}
	}()
	return _d._base.GetJoined(ctx, userId)
}

// GetMessage implements repository.Chat
func (_d ChatWithLog) GetMessage(ctx context.Context, messageId uint) (cp1 *entity.ChatMessage, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("messageId", messageId).
		Msg("store.Chat.GetMessage: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.GetMessage: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetMessage: finished")
		}
	}()
	return _d._base.GetMessage(ctx, messageId)
}

// GetMessages implements repository.Chat
func (_d ChatWithLog) GetMessages(ctx context.Context, userId uint, since uint, limit uint) (cap1 *[]entity.ChatMessage, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("since", since).
		Interface("limit", limit).
		Msg("store.Chat.GetMessages: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.GetMessages: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetMessages: finished")
		}
	}()
	return _d._base.GetMessages(ctx, userId, since, limit)
}

// GetPublic implements repository.Chat
func (_d ChatWithLog) GetPublic(ctx context.Context) (cap1 *[]entity.Channel, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Msg("store.Chat.GetPublic: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.GetPublic: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetPublic: finished")
		}
	}()
	return _d._base.GetPublic(ctx)
}

// Join implements repository.Chat
func (_d ChatWithLog) Join(ctx context.Context, userId uint, channelId uint) (cp1 *entity.Channel, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("channelId", channelId).
		Msg("store.Chat.Join: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.Join: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.Join: finished")
		}
	}()
	return _d._base.Join(ctx, userId, channelId)
}

// Leave implements repository.Chat
func (_d ChatWithLog) Leave(ctx context.Context, userId uint, channelId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("channelId", channelId).
		Msg("store.Chat.Leave: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.Leave: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.Leave: finished")
		}
	}()
	return _d._base.Leave(ctx, userId, channelId)
}

// ReadMessage implements repository.Chat
func (_d ChatWithLog) ReadMessage() {
	log.Trace().Msg("store.Chat.ReadMessage: calling")
	defer func() {
		log.Trace().
			Msg("store.Chat.ReadMessage: finished")
	}()
	_d._base.ReadMessage()
	return
}

// SendMessage implements repository.Chat
func (_d ChatWithLog) SendMessage(ctx context.Context, userId uint, channelId uint, content string, isAction bool) (cp1 *entity.ChatMessage, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("channelId", channelId).
		Interface("content", content).
		Interface("isAction", isAction).
		Msg("store.Chat.SendMessage: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Chat.SendMessage: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.SendMessage: finished")
		}
	}()
	return _d._base.SendMessage(ctx, userId, channelId, content, isAction)
}
