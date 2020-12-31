package sql

import (
	"context"
	"github.com/lib/pq"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ChatStore struct {
	SqlStore
}

func newSqlChatStore(sqlStore SqlStore) store.Chat {
	return &ChatStore{sqlStore}
}

func (c ChatStore) Get(ctx context.Context, channelId uint) (*entity.Channel, error) {
	var channel entity.Channel

	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("id = ?", channelId).
		First(&channel).
		Error

	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (c ChatStore) CreatePm(ctx context.Context, userId, targetId uint) (*entity.Channel, error) {
	channel := entity.Channel{
		Name:        "PM",
		Description: "-",
		Type:        entity.ChannelPMType,
		Users:       pq.Int64Array{int64(userId), int64(targetId)},
		ActiveUsers: pq.Int64Array{int64(userId), int64(targetId)},
	}

	err := c.GetMaster().Transaction(func(tx *gorm.DB) error {
		return c.GetMaster().
			WithContext(ctx).
			Table("channels").
			Create(&channel).
			Error
	})
	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (c ChatStore) GetPublic(ctx context.Context) (*[]entity.Channel, error) {
	var channels []entity.Channel

	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("type = ?", entity.ChannelPublicType).
		Find(&channels).
		Error

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (c ChatStore) GetJoined(ctx context.Context, userId uint) (*[]entity.Channel, error) {
	var channels []entity.Channel

	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("active_users @> ARRAY[?]::int[]", userId).
		Find(&channels).
		Error

	if err != nil {
		return nil, err
	}

	return &channels, nil
}

func (c ChatStore) GetMessages(ctx context.Context, userId, since, limit uint) (*[]entity.ChatMessage, error) {
	var msgs []entity.ChatMessage

	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("id > ? AND c.active_users @> ARRAY[?]::int[]", since, userId).
		Preload(clause.Associations).
		Limit(int(limit)).
		Find(&msgs).
		Error

	if err != nil {
		return nil, err
	}

	return &msgs, nil
}

func (c ChatStore) GetUpdates(ctx context.Context, userId, since, channelId, limit uint) (*entity.ChannelUpdates, error) {
	channels, err := c.Chat().GetJoined(ctx, userId)
	if err != nil {
		return nil, err
	}

	messages, err := c.Chat().GetMessages(ctx, userId, since, limit)
	if err != nil {
		return nil, err
	}

	return &entity.ChannelUpdates{
		Presence: channels,
		Messages: messages,
	}, nil
}

func (c ChatStore) SendMessage(ctx context.Context, userId, channelId uint, content string, isAction bool) (*entity.ChatMessage, error) {
	msg := entity.ChatMessage{
		SenderId:  userId,
		ChannelId: channelId,
		Timestamp: time.Now(),
		Content:   content,
		IsAction:  isAction,
	}

	err := c.GetMaster().Transaction(func(tx *gorm.DB) error {
		return c.GetMaster().
			WithContext(ctx).
			Table("message").
			Create(&msg).
			Error
	})
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (c ChatStore) GetMessage(ctx context.Context, messageId uint) (*entity.ChatMessage, error) {
	var msg entity.ChatMessage

	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("id", messageId).
		Preload(clause.Associations).
		First(&msg).
		Error

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (c ChatStore) Join(ctx context.Context, userId, channelId uint) (*entity.Channel, error) {
	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("id = ?", channelId).
		Update("users", gorm.Expr("array_append(array_remove(active_users, ?), ?)", userId, userId)).
		Update(
			"active_users",
			gorm.Expr("array_append(array_remove(active_users, ?), ?)",
				userId,
				userId,
			),
		).
		Error
	if err != nil {
		return nil, err
	}

	return c.Chat().Get(ctx, channelId)
}

func (c ChatStore) Leave(ctx context.Context, userId, channelId uint) error {
	err := c.GetMaster().
		WithContext(ctx).
		Table("channels").
		Where("id = ?", channelId).
		Update("active_users", gorm.Expr(
			"array_remove(active_users, ?)",
			userId,
		)).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (c ChatStore) ReadMessage() {
	panic("implement me")
}
