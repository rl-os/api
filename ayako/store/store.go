package store

//go:generate mockgen -destination=./mocks/generated.go -source=store.go
//go:generate gowrap gen -g -p . -i OAuth -t layers/log.tmpl -o layers/log_oauth.go
//go:generate gowrap gen -g -p . -i Beatmap -t layers/log.tmpl -o layers/log_beatmap.go
//go:generate gowrap gen -g -p . -i BeatmapSet -t layers/log.tmpl -o layers/log_beatmapset.go
//go:generate gowrap gen -g -p . -i User -t layers/log.tmpl -o layers/log_user.go

import (
	"context"
	"github.com/deissh/osu-lazer/ayako/entity"
	"time"
)

type Store interface {
	OAuth() OAuth
	Beatmap() Beatmap
	BeatmapSet() BeatmapSet
	User() User
}

type OAuth interface {
	CreateClient(ctx context.Context, name string, redirect string) (*entity.OAuthClient, error)
	GetClient(ctx context.Context, id uint, secret string) (*entity.OAuthClient, error)

	CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (*entity.OAuthToken, error)
	RevokeToken(ctx context.Context, userId uint, accessToken string) error
	RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error)
	ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error)
}

type Beatmap interface {
	Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error)
	GetBySetId(ctx context.Context, beatmapsetId uint) []entity.Beatmap

	Create(ctx context.Context, from interface{}) (*entity.Beatmap, error)
	CreateBatch(ctx context.Context, from interface{}) (*[]entity.Beatmap, error)
	Update(ctx context.Context, id uint, from interface{}) (*entity.Beatmap, error)
	Delete(ctx context.Context, id uint) error
}

type BeatmapSet interface {
	Get(ctx context.Context, id uint) (*entity.BeatmapSetFull, error)
	GetAll(ctx context.Context, page int, limit int) (*[]entity.BeatmapSet, error)
	ComputeFields(ctx context.Context, set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error)
	SetFavourite(ctx context.Context, userId uint, id uint) (uint, error)
	SetUnFavourite(ctx context.Context, userId uint, id uint) (uint, error)

	GetLatestId(ctx context.Context) (uint, error)
	GetIdsForUpdate(ctx context.Context, limit int) ([]uint, error)
	Create(ctx context.Context, from interface{}) (*entity.BeatmapSetFull, error)
	Update(ctx context.Context, id uint, from interface{}) (*entity.BeatmapSetFull, error)
	Delete(ctx context.Context, id uint) error

	FetchFromBancho(ctx context.Context, id uint) (*entity.BeatmapSetFull, error)
}

type User interface {
	Get(ctx context.Context, userId uint, mode string) (*entity.User, error)
	GetShort(ctx context.Context, userId uint, mode string) (*entity.UserShort, error)
	GetByBasic(ctx context.Context, login, pwd string) (*entity.UserShort, error)
	ComputeFields(ctx context.Context, user entity.User) (*entity.User, error)
	Create(ctx context.Context, name, email, pwd string) (*entity.User, error)
	Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error)

	Activate(ctx context.Context, userId uint) error
	Deactivate(ctx context.Context, userId uint) error
	Ban(ctx context.Context, userId uint, time time.Duration) error
	UnBan(ctx context.Context, userId uint) error
	Mute(ctx context.Context, userId uint, time time.Duration) error
	UnMute(ctx context.Context, userId uint) error

	UpdateLastVisit(ctx context.Context, userId uint) error
}
