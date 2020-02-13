package entity

import (
	"errors"
	"github.com/deissh/osu-lazer/api/pkg/common/utils"
	"github.com/lib/pq"
	"time"
)

var allowedTypes = []string{"PUBLIC", "PRIVATE", "MULTIPLAYER", "SPECTATOR", "TEMPORARY", "PM", "GROUP"}

// Channel data struct
type Channel struct {
	ID          uint             `json:"channel_id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Description string           `json:"description" db:"description"`
	Type        string           `json:"type" db:"type"`
	Icon        utils.NullString `json:"icon,omitempty" db:"icon"`
	CreatedAt   time.Time        `json:"-" db:"created_at"`

	Users       pq.Int64Array `json:"users,omitempty" db:"users"`
	ActiveUsers pq.Int64Array `json:"-" db:"active_users"`
}

func (c *Channel) Check() error {
	if !utils.ContainsString(allowedTypes, c.Type) {
		return errors.New("not allowed channel type")
	}

	return nil
}
