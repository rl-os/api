package app

import (
	"context"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/repository"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

var (
	modes           = []string{"std", "mania", "catch", "taiko"}
	ErrNotFoundUser = errors.New("not_found_user", http.StatusNotFound, "Not found")
)

type UserUseCase struct {
	*App
	UserRepository repository.User
}

func NewUserUseCase(app *App, rep repository.User) *UserUseCase {
	return &UserUseCase{app, rep}
}

// GetUser from repository and return 404 error if not exist
func (a *UserUseCase) GetUser(ctx context.Context, userID uint, mode string) (*entity.User, error) {
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	data, err := a.UserRepository.Get(ctx, userID, mode)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	return data, nil
}

func (a *UserUseCase) UpdateLastVisit(ctx context.Context, userId uint) error {
	err := a.UserRepository.UpdateLastVisit(ctx, userId)
	if err != nil {
		return ErrNotFoundUser.WithCause(err)
	}

	return nil
}
