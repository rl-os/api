package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/store"
)

type ChatStore struct {
	SqlStore
}

func newSqlChatStore(sqlStore SqlStore) store.Chat {
	return &ChatStore{sqlStore}
}

func (c ChatStore) CreatePM(ctx context.Context, userId, targetId uint, message string, isAction bool) (*entity.ChannelNewPm, error) {
	channel, err := c.Chat().GetOrCreatePm(ctx, userId, targetId)
	if err != nil {
		return nil, err
	}

	msg, err := c.Chat().SendMessage(ctx, userId, channel.ID, message, isAction)
	if err != nil {
		return nil, err
	}

	presence, err := c.Chat().GetJoined(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &entity.ChannelNewPm{
		Id:       channel.ID,
		Presence: *presence,
		Messages: *msg,
	}, nil
}

func (c ChatStore) Get(ctx context.Context, channelId uint) (*entity.Channel, error) {
	var channel entity.Channel

	err := c.GetMaster().GetContext(
		ctx,
		&channel,
		`SELECT
			channels.id, channels.name, channels.description,
			channels.type, channels.icon, channels.active_users, channels.users
		FROM channels
		WHERE channels.id = $1;`,
		channelId,
	)
	if err != nil {
		return nil, errors.WithCause("chat_channel_get", 404, "Channel not found", err)
	}

	return &channel, nil
}

func (c ChatStore) GetOrCreatePm(ctx context.Context, userId, targetId uint) (*entity.Channel, error) {
	var channel entity.Channel

	// TODO: return old chat if exist

	err := c.GetMaster().GetContext(
		ctx,
		&channel,
		`INSERT INTO channels (name, description, type, icon, users, active_users)
		SELECT
			name.str, '-', 'PM', 'PM', ARRAY[$1, $2]::int[], ARRAY[$1, $2]::int[]
		FROM (
			select string_agg(users.username, ' & ') as str
			from users
			where users.id in ($1, $2)
		) as name
		RETURNING *`,
		userId,
		targetId,
	)
	if err != nil {
		return nil, errors.WithCause("chat_pm_channel_get", 400, "Channel not created or found", err)
	}

	return &channel, nil
}

func (c ChatStore) GetPublic(ctx context.Context) (*[]entity.Channel, error) {
	var channels []entity.Channel

	err := c.GetMaster().SelectContext(
		ctx,
		&channels,
		`SELECT
			channels.id, channels.name, channels.description,
			channels.type, channels.icon, channels.active_users, channels.users
		FROM channels
		WHERE channels.type = 'PUBLIC'
		ORDER BY channels.created_at;`,
	)
	if err != nil {
		return nil, errors.WithCause("chat_channel_get", 404, "Channels not found", err)
	}

	return &channels, nil
}

func (c ChatStore) GetJoined(ctx context.Context, userId uint) (*[]entity.Channel, error) {
	var channels []entity.Channel

	err := c.GetMaster().SelectContext(
		ctx,
		&channels,
		`SELECT
			channels.id, channels.name, channels.description,
			channels.type, channels.icon, channels.active_users, channels.users
		FROM channels
		WHERE channels.active_users @> ARRAY[$1]::int[]
		ORDER BY channels.created_at;`,
		userId,
	)
	if err != nil {
		return nil, errors.WithCause("chat_channel_get", 404, "Channels not found", err)
	}

	return &channels, nil
}

func (c ChatStore) GetMessages(ctx context.Context, userId, since uint) (*[]entity.ChatMessage, error) {
	messages := make([]entity.ChatMessage, 0)

	err := c.GetMaster().SelectContext(
		ctx,
		&messages,
		`SELECT
			message.id, message.sender_id, message.channel_id,
			message.created_at, message.content, message.is_action,
			json_build_object('id', u.id, 'username', u.username, 'avatar_url', u.avatar_url,
							  'country_code', u.country_code, 'is_active', u.is_active, 'is_bot', u.is_bot,
							  'is_supporter', u.is_supporter, 'is_online', check_online(u.last_visit)) as sender
		FROM message
				 INNER JOIN users u on message.sender_id = u.id
				 INNER JOIN channels c on message.channel_id = c.id
		WHERE message.id > $2 AND c.active_users @> ARRAY[$1]::int[]
		GROUP BY message.id, u.id
		ORDER BY message.id`,
		userId,
		since,
	)
	if err != nil {
		return nil, errors.WithCause("chat_messages", 404, "Messages not found", err)
	}

	return &messages, nil
}

func (c ChatStore) GetUpdates(ctx context.Context, userId, since, channelId, limit uint) (*entity.ChannelUpdates, error) {
	var updates entity.ChannelUpdates
	if limit <= 0 || limit > 50 {
		limit = 25
	}

	if channelId == 0 {
		channels, err := c.Chat().GetJoined(ctx, userId)
		if err != nil {
			return nil, err
		}

		updates.Presence = *channels
	}

	messages, err := c.Chat().GetMessages(ctx, userId, since)
	if err != nil {
		return nil, err
	}
	updates.Messages = *messages

	return &updates, nil
}

func (c ChatStore) SendMessage(ctx context.Context, userId, channelId uint, content string, isAction bool) (*entity.ChatMessage, error) {
	var message entity.ChatMessage

	err := c.GetMaster().GetContext(
		ctx,
		&message,
		`INSERT INTO message (sender_id, channel_id, content, is_action)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
		userId, channelId, content, isAction,
	)
	if err != nil {
		return nil, errors.WithCause("chat_message_send", 400, "Message not created", err)
	}

	return c.Chat().GetMessage(ctx, message.ID)
}

func (c ChatStore) GetMessage(ctx context.Context, messageId uint) (*entity.ChatMessage, error) {
	var message entity.ChatMessage

	err := c.GetMaster().GetContext(
		ctx,
		&message,
		`SELECT
			message.id, message.sender_id, message.channel_id,
			message.created_at, message.content, message.is_action,
			json_build_object('id', u.id, 'username', u.username, 'avatar_url', u.avatar_url,
			    'country_code', u.country_code, 'is_active', u.is_active, 'is_bot', u.is_bot,
			    'is_supporter', u.is_supporter, 'is_online', check_online(u.last_visit)) as sender
		FROM message
		INNER JOIN users u on message.sender_id = u.id
		WHERE message.id = $1`,
		messageId,
	)
	if err != nil {
		return nil, errors.WithCause("chat_message_send", 400, "Message not created", err)
	}

	return &message, nil
}

func (c ChatStore) Join(ctx context.Context, userId, channelId uint) (*entity.Channel, error) {
	_, err := c.GetMaster().ExecContext(
		ctx,
		`UPDATE channels
				SET users = array_append(array_remove(users, $1), $1),
				    active_users = array_append(array_remove(active_users, $1), $1)
				WHERE channels.id = $2;`,
		userId,
		channelId,
	)
	if err != nil {
		return nil, errors.WithCause("chat_channel_join", 404, "Channel not found", err)
	}

	return c.Chat().Get(ctx, channelId)
}

func (c ChatStore) Leave(ctx context.Context, userId, channelId uint) error {
	_, err := c.GetMaster().ExecContext(
		ctx,
		`UPDATE channels
				SET active_users = array_remove(active_users, $1)
				WHERE channels.id = $2;`,
		userId,
		channelId,
	)
	if err != nil {
		return errors.WithCause("chat_channel_leave", 400, "Channel not found or something else", err)
	}

	return nil
}

func (c ChatStore) ReadMessage() {
	panic("implement me")
}
