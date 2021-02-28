package app

import (
	"context"
	"github.com/deissh/go-utils"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

var (
	modes           = []string{"std", "mania", "catch", "taiko"}
	ErrNotFoundUser = errors.New("not_found_user", http.StatusNotFound, "Not found")
)

type User struct {
	*App
}

// GetUser from store and return 404 error if not exist
func (a *User) Get(ctx context.Context, userID uint, mode string) (*entity.User, error) {
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	data, err := a.Store.User().Get(ctx, userID, mode)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	return data, nil
}
