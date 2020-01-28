package entity

// ChannelUpdates data struct
type ChannelUpdates struct {
	Presence []Channel     `json:"presence"`
	Messages []ChatMessage `json:"messages"`
}
