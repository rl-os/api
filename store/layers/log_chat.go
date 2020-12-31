package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"github.com/rs/zerolog/log"
)

var chatDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "app_store_chat_duration_seconds",
		Help:       "chat runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// ChatWithLog implements store.Chat that is instrumented with zerolog
type ChatWithLog struct {
	_base store.Chat
}

func NewChatWithLog(base store.Chat) store.Chat {
	return ChatWithLog{
		_base: base,
	}
}

// CreatePm implements store.Chat
func (_d ChatWithLog) CreatePm(ctx context.Context, userId uint, targetId uint) (cp1 *entity.Channel, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Chat.CreatePm: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.CreatePm: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.CreatePm: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "CreatePm", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.CreatePm(ctx, userId, targetId)
}

// Get implements store.Chat
func (_d ChatWithLog) Get(ctx context.Context, channelId uint) (cp1 *entity.Channel, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("channelId", channelId).
		Msg("store.Chat.Get: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.Get: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "Get", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Get(ctx, channelId)
}

// GetJoined implements store.Chat
func (_d ChatWithLog) GetJoined(ctx context.Context, userId uint) (cap1 *[]entity.Channel, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Chat.GetJoined: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.GetJoined: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetJoined: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "GetJoined", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetJoined(ctx, userId)
}

// GetMessage implements store.Chat
func (_d ChatWithLog) GetMessage(ctx context.Context, messageId uint) (cp1 *entity.ChatMessage, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("messageId", messageId).
		Msg("store.Chat.GetMessage: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.GetMessage: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetMessage: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "GetMessage", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetMessage(ctx, messageId)
}

// GetMessages implements store.Chat
func (_d ChatWithLog) GetMessages(ctx context.Context, userId uint, since uint, limit uint) (cap1 *[]entity.ChatMessage, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("since", since).
		Interface("limit", limit).
		Msg("store.Chat.GetMessages: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.GetMessages: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetMessages: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "GetMessages", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetMessages(ctx, userId, since, limit)
}

// GetPublic implements store.Chat
func (_d ChatWithLog) GetPublic(ctx context.Context) (cap1 *[]entity.Channel, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Msg("store.Chat.GetPublic: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.GetPublic: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetPublic: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "GetPublic", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetPublic(ctx)
}

// GetUpdates implements store.Chat
func (_d ChatWithLog) GetUpdates(ctx context.Context, userId uint, since uint, channelId uint, limit uint) (cp1 *entity.ChannelUpdates, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("since", since).
		Interface("channelId", channelId).
		Interface("limit", limit).
		Msg("store.Chat.GetUpdates: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.GetUpdates: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.GetUpdates: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "GetUpdates", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetUpdates(ctx, userId, since, channelId, limit)
}

// Join implements store.Chat
func (_d ChatWithLog) Join(ctx context.Context, userId uint, channelId uint) (cp1 *entity.Channel, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("channelId", channelId).
		Msg("store.Chat.Join: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.Join: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.Join: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "Join", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Join(ctx, userId, channelId)
}

// Leave implements store.Chat
func (_d ChatWithLog) Leave(ctx context.Context, userId uint, channelId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("channelId", channelId).
		Msg("store.Chat.Leave: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.Leave: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.Leave: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "Leave", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Leave(ctx, userId, channelId)
}

// ReadMessage implements store.Chat
func (_d ChatWithLog) ReadMessage() {
	_since := time.Now()
	log.Trace().Msg("store.Chat.ReadMessage: calling")
	defer func() {
		result := "ok"
		log.Trace().
			Msg("store.Chat.ReadMessage: finished")
		chatDurationSummaryVec.WithLabelValues("Chat", "ReadMessage", result).Observe(time.Since(_since).Seconds())
	}()
	_d._base.ReadMessage()
	return
}

// SendMessage implements store.Chat
func (_d ChatWithLog) SendMessage(ctx context.Context, userId uint, channelId uint, content string, isAction bool) (cp1 *entity.ChatMessage, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("channelId", channelId).
		Interface("content", content).
		Interface("isAction", isAction).
		Msg("store.Chat.SendMessage: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Chat.SendMessage: returned an error")
		} else {
			log.Trace().
				Msg("store.Chat.SendMessage: finished")
		}
		chatDurationSummaryVec.WithLabelValues("Chat", "SendMessage", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.SendMessage(ctx, userId, channelId, content, isAction)
}
