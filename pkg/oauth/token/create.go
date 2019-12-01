package token

import (
	"github.com/labstack/echo/v4"
)

// CreateTokenRequestData contain incoming data with user credentials
type CreateTokenRequestData struct {
	GrantType    string `json:"grant_type" form:"grant_type" validate:"required"`
	ClientID     string `json:"client_id" from:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" from:"client_secret" validate:"required"`
	Scope        string `json:"scope" form:"scope" validate:"required"`

	// grant_type == password
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`

	// grant_type == refresh_token
	RefreshToken string `json:"refresh_token" from:"refresh_token"`
}

// CreateTokenResponseData response struct that contain new access_token and refresh_token
type CreateTokenResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// CreateTokenHandler create new access_token and refresh_token pare
func CreateTokenHandler(c echo.Context) (err error) {
	return nil
}
