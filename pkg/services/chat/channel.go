package chat

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetChannels of all joinable public channels
func GetChannel(id uint) (*entity.Channel, error) {
	var channel entity.Channel

	err := pkg.Db.Get(
		&channel,
		`SELECT channels.id, name, description, type, icon
				FROM channels
				WHERE id = $1`,
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
		`SELECT channels.id, name, description, type, icon
				FROM channels
				INNER JOIN user_channels uc on channels.id = uc.channel_id
				WHERE type = 'PUBLIC' AND user_id = $1`,
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
