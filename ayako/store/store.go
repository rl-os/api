package store

//go:generate mockgen -destination=./mocks/generated.go -source=store.go
//go:generate gowrap gen -g -p . -i Beatmap -t layers/log.tmpl -o layers/log_beatmap.go
//go:generate gowrap gen -g -p . -i BeatmapSet -t layers/log.tmpl -o layers/log_beatmapset.go

import (
	"github.com/deissh/osu-lazer/ayako/entity"
)

type Store interface {
	Beatmap() Beatmap
	BeatmapSet() BeatmapSet
}

type Beatmap interface {
	Get(id uint) (*entity.SingleBeatmap, error)
	GetBySetId(beatmapsetId uint) []entity.Beatmap

	Create(from interface{}) (*entity.Beatmap, error)
	CreateBatch(from interface{}) (*[]entity.Beatmap, error)
	Update(id uint, from interface{}) (*entity.Beatmap, error)
	Delete(id uint) error
}

type BeatmapSet interface {
	Get(id uint) (*entity.BeatmapSetFull, error)
	GetAll(page int, limit int) (*[]entity.BeatmapSet, error)
	ComputeFields(set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error)

	GetLatestId() (uint, error)
	GetIdsForUpdate(limit int) ([]uint, error)
	Create(from interface{}) (*entity.BeatmapSetFull, error)
	Update(id uint, from interface{}) (*entity.BeatmapSetFull, error)
	Delete(id uint) error

	FetchFromBancho(id uint) (*entity.BeatmapSetFull, error)
}
