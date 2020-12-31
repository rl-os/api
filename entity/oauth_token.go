package entity

import "time"

// OAuthToken contains access_token and refresh
type OAuthToken struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ClientID     uint      `json:"client_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Scopes       string    `json:"scopes"`
	Revoked      bool      `json:"revoked"`
	ExpiresIn    int       `json:"expires_in" gorm:"-"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}
