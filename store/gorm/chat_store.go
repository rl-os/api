package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"gorm.io/gorm"
)

type ChatStore struct {
	SqlStore
}

func newSqlChatStore(sqlStore SqlStore) store.Chat {
	return &ChatStore{sqlStore}
}

func (c ChatStore) CreatePM(ctx context.Context, userId, targetId uint, message string, isAction bool) (*entity.ChannelNewPm, error) {
	panic("implement me")
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

func (c ChatStore) GetOrCreatePm(ctx context.Context, userId, targetId uint) (*entity.Channel, error) {
	panic("implement me")
}

func (c ChatStore) GetPublic(ctx context.Context) (*[]entity.Channel, error) {
	panic("implement me")
}

func (c ChatStore) GetJoined(ctx context.Context, userId uint) (*[]entity.Channel, error) {
	panic("implement me")
}

func (c ChatStore) GetMessages(ctx context.Context, userId, since uint) (*[]entity.ChatMessage, error) {
	panic("implement me")
}

func (c ChatStore) GetUpdates(ctx context.Context, userId, since, channelId, limit uint) (*entity.ChannelUpdates, error) {
	panic("implement me")
}

func (c ChatStore) SendMessage(ctx context.Context, userId, channelId uint, content string, isAction bool) (*entity.ChatMessage, error) {
	panic("implement me")
}

func (c ChatStore) GetMessage(ctx context.Context, messageId uint) (*entity.ChatMessage, error) {
	panic("implement me")
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
