package chat

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetChannels of all joinable public channels
func GetChannels() (*[]entity.Channel, error) {
	var defaultChannels []entity.Channel

	err := pkg.Db.Select(
		&defaultChannels,
		`SELECT channels.id, channels.name, channels.description, channels.type, channels.icon
				FROM channels
				WHERE channels.type = 'PUBLIC';`,
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
