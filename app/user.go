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
	ErrNotFoundUser = errors.New("user", http.StatusNotFound, "User not found")
)

// GetUser from store and return 404 error if not exist
func (a *App) GetUser(ctx context.Context, userID uint, mode string) (*entity.User, error) {
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	data, err := a.Store.User().Get(ctx, userID, mode)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	return data, nil
}
