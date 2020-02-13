package oauth

import (
	"fmt"
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gookit/config/v2"
	"github.com/labstack/echo/v4"
	"time"
)

// Token model
type Token struct {
	ID           uint   `db:"id" json:"id"`
	UserID       uint   `db:"user_id" json:"user_id"`
	ClientID     uint   `db:"client_id" json:"client_id"`
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
	Scopes       string `db:"scopes" json:"scopes"`
	Revoked      bool   `db:"revoked" json:"revoked" json:"revoked"`
	ExpiresIn    int    `json:"expires_in"`

	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// CreateOAuthToken and return model with valid access_token and refresh_token
func CreateOAuthToken(userID uint, clientID uint, clientSecret string, scopes string) (Token, error) {
	_, err := FindOAuthClient(clientID, clientSecret)
	if err != nil {
		return Token{}, echo.NewHTTPError(400, "Not founded client_id or client_secret.")
	}

	// generate random refresh_token
	refreshToken, err := utils.GenerateRandomString(255)
	jwtID, err := utils.GenerateRandomString(64)
	if err != nil {
		return Token{}, echo.NewHTTPError(500, "New refresh token generate error.")
	}

	now := time.Now()
	expireAt := now.Add(
		time.Hour * time.Duration(config.Int("server.jwt.hours_before_revoke", 1)),
	)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"aud":    fmt.Sprint(clientID),
		"jti":    jwtID,
		"iat":    now.Unix(),
		"exp":    expireAt.Unix(),
		"sub":    fmt.Sprint(userID),
		"scopes": []string{scopes},
	})
	accessToken, err := tokenClaims.SignedString([]byte(config.String("server.jwt.secret")))
	if err != nil {
		return Token{}, echo.NewHTTPError(500, "New access token generate error.")
	}

	// inserting new oauth_token
	token := Token{
		UserID:       userID,
		ClientID:     clientID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Scopes:       scopes,
		ExpiresIn:    int(expireAt.Sub(now).Seconds()),
		ExpiresAt:    expireAt,
	}
	err = pkg.Db.Get(
		&token,
		`INSERT INTO oauth_token (user_id, client_id, access_token, refresh_token, scopes, expires_at)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING *`,
		token.UserID, token.ClientID, token.AccessToken, token.RefreshToken, token.Scopes, token.ExpiresAt,
	)
	if err != nil {
		return Token{}, echo.NewHTTPError(500, "Creating new access_token in database error.")
	}

	return token, nil
}

// RevokeOAuthToken and return error if not successfully
func RevokeOAuthToken() (err error) {
	return nil
}

// ValidateOAuthToken and return OAuthToken with full information
func ValidateOAuthToken(accessToken string) (Token, error) {
	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.String("server.jwt.secret")), nil
	})

	v, _ := err.(*jwt.ValidationError)
	if err != nil && v.Errors == jwt.ValidationErrorExpired {
		_, _ = pkg.Db.Exec(`UPDATE oauth_token SET revoked = true WHERE access_token = $1`, accessToken)

		return Token{}, pkg.NewHTTPError(401, "oauth_token_revoked", "Access token expired")
	} else if err != nil {
		return Token{}, pkg.NewHTTPError(401, "oauth_token_error", "Invalid access token")
	}

	var token Token
	err = pkg.Db.Get(
		&token,
		`SELECT * FROM oauth_token
				WHERE access_token = $1`,
		accessToken,
	)
	if err != nil {
		return Token{}, echo.NewHTTPError(500, "Selecting access_token in database error.")
	}
	if token.Revoked {
		return Token{}, pkg.NewHTTPError(401, "oauth_token_revoked", "Access token expired")
	}

	return token, nil
}

// RefreshOAuthToken create new access_token and return it
func RefreshOAuthToken(refreshToken string, clientID uint, clientSecret string, scopes string) (Token, error) {
	var token Token
	err := pkg.Db.Get(
		&token,
		`UPDATE oauth_token
			SET revoked = true
			WHERE refresh_token = $1 AND revoked = false
			RETURNING *`,
		refreshToken,
	)
	if err != nil {
		return Token{}, pkg.NewHTTPError(400, "oauth_token_not_exist", "Refresh token not exist or already revoked")
	}

	token, err = CreateOAuthToken(token.UserID, clientID, clientSecret, scopes)
	return token, err
}
