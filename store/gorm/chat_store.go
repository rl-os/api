package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
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
	panic("implement me")
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
	panic("implement me")
}

func (c ChatStore) Leave(ctx context.Context, userId, channelId uint) error {
	panic("implement me")
}

func (c ChatStore) ReadMessage() {
	panic("implement me")
}
