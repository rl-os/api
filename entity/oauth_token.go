package entity

import "time"

// OAuthToken contains access_token and refresh
type OAuthToken struct {
	ID           uint      `json:"id"`
	UserID       uint      `db:"user_id" json:"user_id"`
	ClientID     uint      `db:"client_id" json:"client_id"`
	AccessToken  string    `db:"access_token" json:"access_token"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
	Scopes       string    `json:"scopes"`
	Revoked      bool      `json:"revoked"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
