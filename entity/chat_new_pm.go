package entity

type ChannelNewPm struct {
	Id       uint        `json:"new_channel_id"`
	Presence []Channel   `json:"presence"`
	Messages ChatMessage `json:"messages"`
}
