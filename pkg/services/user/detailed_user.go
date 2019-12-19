package user

import (
	"github.com/deissh/osu-api-server/pkg"
	"net/http"
)

type DetailedUser struct {
	BaseUser
	Detail
}

// GetDetailedUser and compute some fields
func GetDetailedUser(id uint) (DetailedUser, error) {
	var user DetailedUser

	err := pkg.Db.Get(
		&user,
		`SELECT * FROM users
    			INNER JOIN user_details ON users.id = user_details.user_id
				WHERE users.id = $1`,
		id,
	)
	if err != nil {
		return DetailedUser{}, pkg.NewHTTPError(http.StatusUnauthorized, "user_not_founded", "User not founded.")
	}

	user.Compute()

	return DetailedUser{}, nil
}
