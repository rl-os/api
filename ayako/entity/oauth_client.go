package entity

import "time"

// OAuthClient of oauth2 server
type OAuthClient struct {
	ID        uint      `json:"id"`
	UserID    uint      `db:"user_id" json:"user_id"`
	Name      string    `json:"name"`
	Secret    string    `json:"secret"`
	Redirect  string    `json:"redirect"`
	Revoked   string    `json:"revoked"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
