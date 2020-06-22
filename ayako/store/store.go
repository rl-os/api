package store

//go:generate mockgen -destination=./mocks/generated.go -source=store.go
//go:generate gowrap gen -g -p . -i OAuth -t layers/log.tmpl -o layers/log_oauth.go
//go:generate gowrap gen -g -p . -i Beatmap -t layers/log.tmpl -o layers/log_beatmap.go
//go:generate gowrap gen -g -p . -i BeatmapSet -t layers/log.tmpl -o layers/log_beatmapset.go

import (
	"context"
	"github.com/deissh/osu-lazer/ayako/entity"
)

type Store interface {
	//OAuth() OAuth
	Beatmap() Beatmap
	BeatmapSet() BeatmapSet
}

type OAuth interface {
	CreateClient(ctx context.Context, name string, redirect string)
	GetClient(ctx context.Context, id uint)

	CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string)
	RevokeToken(ctx context.Context, userId uint, accessToken string)
	RefreshToken(ctx context.Context, refreshToken string)
	ValidateToken(ctx context.Context, accessToken string)
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
