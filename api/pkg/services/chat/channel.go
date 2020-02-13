package chat

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetChannel by id
func GetChannel(id uint) (*entity.Channel, error) {
	var channel entity.Channel

	err := pkg.Db.Get(
		&channel,
		`SELECT
       				channels.id, channels.name, channels.description,
       				channels.type, channels.icon, channels.active_users, channels.users
				FROM channels
				WHERE channels.id = $1;`,
		id,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("public channels not found")
		return nil, pkg.NewHTTPError(http.StatusNotFound, "chat_channels", "Channel not found.")
	}

	return &channel, nil
}

// GetChannels of all joinable public channels
func GetChannels() (*[]entity.Channel, error) {
	var defaultChannels []entity.Channel

	err := pkg.Db.Select(
		&defaultChannels,
		`SELECT
       				channels.id, channels.name, channels.description,
       				channels.type, channels.icon, channels.active_users, channels.users
				FROM channels
				WHERE channels.type = 'PUBLIC'
				ORDER BY channels.created_at;`,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("public channels not found")
		return nil, pkg.NewHTTPError(http.StatusNotFound, "chat_channels", "Channels not found.")
	}

	return &defaultChannels, nil
}

// GetChannels of all user channels
func GetUserChannels(userId uint) (*[]entity.Channel, error) {
	defaultChannels := make([]entity.Channel, 0)

	err := pkg.Db.Select(
		&defaultChannels,
		`SELECT
       				channels.id, channels.name, channels.description,
       				channels.type, channels.icon, channels.active_users, channels.users
				FROM channels
				WHERE channels.active_users @> ARRAY[$1]::int[]
				ORDER BY channels.created_at;`,
		userId,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("public channels not found")
		return nil, pkg.NewHTTPError(http.StatusNotFound, "chat_channels", "Channels not found.")
	}

	return &defaultChannels, nil
}

// GetUpdates from *since* in channel_id
func GetUpdates(userId uint, since uint, channelId uint, limit uint) (*entity.ChannelUpdates, error) {
	var updates entity.ChannelUpdates
	if limit <= 0 || limit > 50 {
		limit = 25
	}

	if channelId == 0 {
		channels, err := GetUserChannels(userId)
		if err != nil {
			return nil, err
		}

		updates.Presence = *channels
	}

	messages, err := GetMessagesAll(userId, since)
	if err != nil {
		return nil, err
	}
	updates.Messages = *messages

	return &updates, nil
}

// Join user to channel
func Join(userId uint, channelId uint) (*entity.Channel, error) {
	_, err := pkg.Db.Exec(
		`UPDATE channels
				SET users = array_append(array_remove(users, $1), $1),
				    active_users = array_append(array_remove(active_users, $1), $1)
				WHERE channels.id = $2;`,
		userId,
		channelId,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Uint("channel_id", channelId).
			Uint("user_id", userId).
			Msg("user not joined to channel")
		return nil, pkg.NewHTTPError(http.StatusBadRequest, "chat_channels", "User not joined to channel.")
	}

	channel, err := GetChannel(channelId)
	return channel, nil
}

// Leave user to channel
func Leave(userId uint, channelId uint) error {
	_, err := pkg.Db.Exec(
		`UPDATE channels
				SET active_users = array_remove(active_users, $1)
				WHERE channels.id = $2;`,
		userId,
		channelId,
	)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "chat_channels", "User not leaved to channel.")
	}

	return nil
}

// GetPm channel or create new if not exist
func GetPm(userId uint, secondId uint) (*entity.Channel, error) {
	var channel entity.Channel

	err := pkg.Db.Get(
		&channel,
		`INSERT INTO channels (name, description, type, icon, users, active_users)
				VALUES ('PM', '-', 'PM', DEFAULT, ARRAY[$1, $2]::int[], ARRAY[$1, $2]::int[])
				RETURNING *`,
		userId,
		secondId,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("pm channels not created")
		return nil, pkg.NewHTTPError(http.StatusBadRequest, "chat_channels", "Pm channels not created.")
	}

	return &channel, err
}

// NewPm between two users
func NewPm(userId uint, targetId uint, content string, isAction bool) (*entity.ChannelNewPm, error) {
	channel, err := GetPm(userId, targetId)
	if err != nil {
		return nil, err
	}

	message, err := SendMessage(userId, channel.ID, content, isAction)
	if err != nil {
		return nil, err
	}

	presence, err := GetUserChannels(userId)
	if err != nil {
		return nil, err
	}

	return &entity.ChannelNewPm{
		Id:       channel.ID,
		Messages: *message,
		Presence: *presence,
	}, nil
}
