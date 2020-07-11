package entity

import (
	"time"
)

type ChatMessage struct {
	ID        uint      `json:"message_id" db:"id"`
	SenderId  uint      `json:"sender_id" db:"sender_id"`
	ChannelId uint      `json:"channel_id" db:"channel_id"`
	Timestamp time.Time `json:"timestamp" db:"created_at"`
	Content   string    `json:"content" db:"content"`
	IsAction  bool      `json:"is_action" db:"is_action"`

	Sender UserShortField `json:"sender"`
}
