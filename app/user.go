package app

import (
	"context"
	"fmt"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/repository"
	"net/http"
)

var (
	modes             = []string{"std", "mania", "catch", "taiko"}
	ErrNotFoundUser   = errors.New("not_found_user", http.StatusNotFound, "Not found")
	ErrUserNotCreated = errors.New("user_not_created", http.StatusBadRequest, "Bad request")
)

type UserUseCase struct {
	*App
	UserRepository repository.User
}

func NewUserUseCase(app *App, rep repository.User) *UserUseCase {
	return &UserUseCase{app, rep}
}

// Get from repository and return 404 error if not exist
func (a *UserUseCase) Get(ctx context.Context, userID uint, mode string) (*entity.User, error) {
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	var cached entity.User

	key := fmt.Sprintf("%d__%s", userID, mode)
	err := a.Cache.Get("user", key, &cached)
	if err == nil {
		return &cached, nil
	}

	data, err := a.UserRepository.Get(ctx, userID, mode)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	a.Cache.Set("user", key, data)

	return data, nil
}

func (a *UserUseCase) UpdateLastVisit(ctx context.Context, userId uint) error {
	err := a.UserRepository.UpdateLastVisit(ctx, userId)
	if err != nil {
		return ErrNotFoundUser.WithCause(err)
	}

	return nil
}

func (a UserUseCase) Create(ctx context.Context, user request.CreateUser) (*entity.User, error) {
	if err := a.Validator.Struct(&user); err != nil {
		return nil, InvalidAuthParamsErr.WithCause(err)
	}

	data, err := a.UserRepository.Create(ctx, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, ErrUserNotCreated.WithCause(err)
	}

	return data, nil
}
