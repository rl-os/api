package oauth

import (
	"github.com/deissh/go-utils"
	"github.com/deissh/osu-lazer/api/pkg"
	"time"
)

// Client model
type Client struct {
	ID        uint      `db:"id" json:"id"`
	UserID    uint      `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Secret    string    `db:"secret" json:"secret"`
	Redirect  string    `db:"redirect" json:"redirect"`
	Revoked   string    `db:"revoked" json:"revoked"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// CreateOAuthClient and return it
func CreateOAuthClient(userID uint, name string, redirect string) (client Client, err error) {
	secret, err := utils.GenerateRandomString(255)
	if err != nil {
		return Client{}, pkg.NewHTTPError(500, "server_error", "New refresh token generate error.")
	}

	err = pkg.Db.Get(
		&client,
		`INSERT INTO oauth_client (user_id, name, secret, redirect)
				VALUES ($1, $2, $3, $4)
				RETURNING *`,
		userID, name, secret, redirect,
	)
	if err != nil {
		return Client{}, pkg.NewHTTPError(500, "server_error", "Creating new oauth_client in database error.")
	}

	return
}

// FindOAuthClient by clientId and secretId
func FindOAuthClient(clientID uint, clientSecret string) (Client, error) {
	out := Client{}

	err := pkg.Db.Get(
		&out,
		`SELECT * FROM oauth_client WHERE id = $1 AND secret = $2`,
		clientID, clientSecret,
	)
	return out, err
}
