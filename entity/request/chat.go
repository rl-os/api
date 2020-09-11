package request

// CreateNewChat contain incoming data
type CreateNewChat struct {
	TargetId uint   `json:"target_id" query:"target_id" form:"target_id"`
	Message  string `json:"message" query:"message" form:"message"`
	IsAction bool   `json:"is_action" query:"is_action" form:"is_action"`
}

// GetChatUpdates contain incoming data
type GetChatUpdates struct {
	Since     uint `json:"since" query:"since"`
	ChannelId uint `json:"channel_id" query:"channel_id"`
	Limit     uint `json:"limit" query:"limit"`
}

// GetMessages contain incoming data
type GetMessages struct {
	Limit uint `json:"limit" query:"limit"`
}

// SendMessage contain incoming data
type SendMessage struct {
	Message  string `json:"message" form:"message" validate:"required"`
	IsAction bool   `json:"is_action" form:"is_action"`
}
