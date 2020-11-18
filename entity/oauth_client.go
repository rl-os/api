package entity

import "time"

// OAuthClient of oauth2 server
type OAuthClient struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Secret    string    `json:"secret"`
	Redirect  string    `json:"redirect"`
	Revoked   bool      `json:"revoked"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName of OAuthClient
func (c OAuthClient) TableName() string {
	return "oauth_client"
}
