package entity

import (
	"github.com/deissh/go-utils"
	"github.com/lib/pq"
	"time"
)

const (
	ChannelPublicType  = "PUBLIC"
	ChannelPrivateType = "PRIVATE"
	ChannelMPType      = "MULTIPLAYER"
	ChannelSPType      = "SPECTATOR"
	ChannelTMPType     = "TEMPORARY"
	ChannelPMType      = "PM"
	ChannelGroupType   = "GROUP"
)

var _ = []string{
	ChannelPublicType, ChannelPrivateType,
	ChannelMPType, ChannelSPType, ChannelTMPType,
	ChannelPMType, ChannelGroupType,
}

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
