package sql

import (
	"fmt"
	"github.com/deissh/go-utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"time"
)

// tokenCreate and fill all fields in struct
func tokenCreate(jwtSecret string, beforeRevoke time.Duration, userID uint, clientID uint, scopes string) (*entity.OAuthToken, error) {
	// generate random refresh_token
	refreshToken, err := utils.GenerateRandomString(255)
	jwtID, err := utils.GenerateRandomString(64)
	if err != nil {
		return nil, errors.WithCause("internal_signing", 500, "new refresh token generate", err)
	}

	now := time.Now()
	expireAt := now.Add(beforeRevoke)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"aud":    fmt.Sprint(clientID),
		"jti":    jwtID,
		"iat":    now.Unix(),
		"exp":    expireAt.Unix(),
		"sub":    fmt.Sprint(userID),
		"scopes": []string{scopes},
	})

	accessToken, err := tokenClaims.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, errors.WithCause("internal_signing", 500, "access token signing", err)
	}

	// inserting new oauth_token
	token := entity.OAuthToken{
		UserID:       userID,
		ClientID:     clientID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Scopes:       scopes,
		ExpiresIn:    int(expireAt.Sub(now).Seconds()),
		ExpiresAt:    expireAt,
	}

	return &token, nil
}
