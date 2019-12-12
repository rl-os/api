package oauth

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/utils"
	"github.com/labstack/echo/v4"
	"time"
)

// OAuthClient model
type OAuthClient struct {
	Id       int    `db:"id" json:"id"`
	UserId   int    `db:"user_id" json:"user_id"`
	Name     string `db:"name" json:"name"`
	Secret   string `db:"secret" json:"secret"`
	Redirect string `db:"redirect" json:"redirect"`
	Revoked  string `db:"revoked" json:"revoked"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// CreateOAuthClient and return it
func CreateOAuthClient(userId int, name string, redirect string) (client OAuthClient, err error) {
	secret, err := utils.GenerateRandomString(255)
	if err != nil {
		return OAuthClient{}, echo.NewHTTPError(500, "New refresh token generate error.")
	}

	err = pkg.Db.Get(
		&client,
		`INSERT INTO oauth_client (user_id, name, secret, redirect)
				VALUES ($1, $2, $3)
				RETURNING *`,
		userId, name, secret, redirect,
	)
	if err != nil {
		return OAuthClient{}, echo.NewHTTPError(500, "Creating new oauth_client in database error.")
	}

	return
}

// FindOAuthClient by clientId and secretId
func FindOAuthClient(clientId string, clientSecret string) (OAuthClient, error) {
	out := OAuthClient{}

	err := pkg.Db.Get(
		&out,
		`SELECT * FROM oauth_client WHERE id = $1 AND secret = $2`,
		clientId, clientSecret,
	)
	return out, err
}
