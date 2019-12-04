package services

import (
	"database/sql"
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/utils"
	"github.com/labstack/echo/v4"
	"time"
)

// OAuthToken model
type OAuthToken struct {
	Id           int            `db:"id" json:"id"`
	UserId       int            `db:"user_id" json:"user_id"`
	ClientId     uint           `db:"client_id" json:"client_id"`
	AccessToken  string         `db:"access_token" json:"access_token"`
	RefreshToken string         `db:"refresh_token" json:"refresh_token"`
	Scopes       sql.NullString `db:"scopes" json:"scopes"`
	Revoked      bool           `db:"revoked" json:"revoked" json:"revoked"`
	ExpiresAt    time.Time      `db:"expires_at" json:"expires_at"`
}

// CreateOAuthToken and return model with valid access_token and refresh_token
func CreateOAuthToken(userId int, clientId string, clientSecret string, scopes string) (OAuthToken, error) {
	client, err := FindOAuthClient(clientId, clientSecret)
	if err != nil {
		return OAuthToken{}, echo.NewHTTPError(400, "Not founded client_id or client_secret.")
	}

	// generate random refresh_token
	refreshToken, err := utils.GenerateRandomString(255)
	if err != nil {
		return OAuthToken{}, echo.NewHTTPError(500, "New refresh token generate error.")
	}

	var token OAuthToken

	// inserting new oauth_token
	err = pkg.Db.Get(
		&token,
		`INSERT INTO oauth_token (user_id, client_id, access_token, refresh_token, scopes)
				VALUES ($1, $2, $3, $4)
				RETURNING *`,
		userId, client.Id, "todo", refreshToken, scopes,
	)
	if err != nil {
		return OAuthToken{}, echo.NewHTTPError(500, "Creating new access_token in database error.")
	}

	return token, nil
}

// RevokeOAuthToken and return error if not successfully
func RevokeOAuthToken() (err error) {
	return nil
}

// ValidateOAuthToken and return OAuthToken with full information
func ValidateOAuthToken(accessToken string) (OAuthToken, err error) {
	return nil, nil
}

// RefreshOAuthToken create new access_token and return it
func RefreshOAuthToken(refreshToken string) (OAuthToken, err error) {
	return nil, nil
}
