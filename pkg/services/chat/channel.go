package chat

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	"github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetChannels of all joinable public channels
func GetChannel(id uint) (*entity.Channel, error) {
	var channel entity.Channel

	err := pkg.Db.Get(
		&channel,
		`SELECT
       				channels.id, channels.name, channels.description,
       				channels.type, channels.icon,
       				array_remove(array_agg(uc.user_id), null) as users
				FROM user_channels AS uc
				FULL OUTER JOIN (SELECT
					cu.channel_id
				  	FROM user_channels AS cu
				  	ORDER BY channel_id DESC
				) as uc2 ON uc.id = uc2.channel_id
				FULL OUTER JOIN channels on uc.channel_id = channels.id
				WHERE channels.id = $1
				GROUP BY channels.id;`,
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
       				channels.type, channels.icon,
       				array_remove(array_agg(uc.user_id), null) as users
				FROM user_channels AS uc
				FULL OUTER JOIN (SELECT
					cu.channel_id
				  	FROM user_channels AS cu
				  	ORDER BY channel_id DESC
				) as uc2 ON uc.id = uc2.channel_id
				FULL OUTER JOIN channels on uc.channel_id = channels.id
				WHERE channels.type = 'PUBLIC'
				GROUP BY channels.id;`,
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
			channels.type, channels.icon,
			array_remove(array_agg(uc.user_id), null) as users
		FROM user_channels AS uc
				 FULL OUTER JOIN (SELECT
									  cu.channel_id
								  FROM user_channels AS cu
								  ORDER BY channel_id DESC
		) as uc2 ON uc.id = uc2.channel_id
				 FULL OUTER JOIN channels on uc.channel_id = channels.id
		WHERE uc.user_id = $1 OR uc.channel_id = channels.id
		GROUP BY channels.id;`,
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
		`INSERT INTO user_channels (user_id, channel_id)
				VALUES ($1, $2)`,
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
		`DELETE FROM channels
    			WHERE id = $1 AND type = 'PM'`,
		channelId,
	)
	if err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "chat_channels", "User not leaved to channel.")
	}

	_, err = pkg.Db.Exec(
		`DELETE FROM user_channels
    			WHERE user_id = $1 AND channel_id = $2`,
		userId,
		channelId,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Uint("channel_id", channelId).
			Uint("user_id", userId).
			Msg("user not leave from channel")
		return pkg.NewHTTPError(http.StatusBadRequest, "chat_channels", "User not leaved to channel.")
	}

	return nil
}

// GetPm channel or create new if not exist
func GetPm(userId uint, secondId uint) (*entity.Channel, error) {
	var channel entity.Channel

	secondUser, err := user.GetUser(secondId, "")
	if err != nil {
		return nil, err
	}

	err = pkg.Db.Get(
		&channel,
		`INSERT INTO channels (name, description, type, icon)
				VALUES ($1, '-', 'PM', DEFAULT)
				RETURNING *`,
		secondUser.Username,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("pm channels not created")
		return nil, pkg.NewHTTPError(http.StatusBadRequest, "chat_channels", "Pm channels not created.")
	}

	_, err = Join(userId, channel.ID)
	if err != nil {
		return nil, err
	}
	_, err = Join(secondId, channel.ID)
	if err != nil {
		return nil, err
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
