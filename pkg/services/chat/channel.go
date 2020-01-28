package chat

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	"github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/rs/zerolog/log"
	"net/http"
)

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
	var defaultChannels []entity.Channel

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

func SendMessage(senderId uint, channelId uint, content string, IsAction bool) (*entity.ChatMessage, error) {
	var message entity.ChatMessage

	err := pkg.Db.Get(
		&message,
		`INSERT INTO message (sender_id, channel_id, content, is_action)
				VALUES ($1, $2, $3, $4)
				RETURNING *`,
		senderId,
		channelId,
		content,
		IsAction,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("message not send")
		return nil, pkg.NewHTTPError(http.StatusBadRequest, "channel_message", "Message not send.")
	}

	res, err := user.GetUser(message.SenderId, "")
	if err != nil {
		return nil, err
	}

	message.Sender = res.GetShort()

	return &message, nil
}