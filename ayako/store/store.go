package store

//go:generate mockgen -destination=./mocks/generated.go -source=store.go

import (
	"github.com/deissh/osu-lazer/ayako/entity"
)

type Store interface {
	Beatmap() Beatmap
	BeatmapSet() BeatmapSet
}

type Beatmap interface {
	GetBeatmap(id uint) (*entity.Beatmap, error)

	CreateBeatmap(from interface{}) (*entity.Beatmap, error)
	UpdateBeatmap(id uint, from interface{}) (*entity.Beatmap, error)
	DeleteBeatmap(id uint) error
}

type BeatmapSet interface {
	GetBeatmapSet(id uint) (*entity.BeatmapSetFull, error)
	GetAllBeatmapSets(page int, limit int) (*[]entity.BeatmapSet, error)
	GetBeatmapSetIdForUpdate(limit int) ([]uint, error)

	GetLatestBeatmapId() (uint, error)
	CreateBeatmapSet(from interface{}) (*entity.BeatmapSetFull, error)
	UpdateBeatmapSet(id uint, from interface{}) (*entity.BeatmapSetFull, error)
	DeleteBeatmapSet(id uint) error

	Fetch(id uint) (*entity.BeatmapSetFull, error)
}
