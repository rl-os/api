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
	GetBeatmap(id uint) (*entity.SingleBeatmap, error)
	GetBeatmapsBySet(beatmapsetId uint) []entity.Beatmap

	CreateBeatmap(from interface{}) (*entity.Beatmap, error)
	CreateBeatmaps(from interface{}) (*[]entity.Beatmap, error)
	UpdateBeatmap(id uint, from interface{}) (*entity.Beatmap, error)
	DeleteBeatmap(id uint) error
}

type BeatmapSet interface {
	GetBeatmapSet(id uint) (*entity.BeatmapSetFull, error)
	GetAllBeatmapSets(page int, limit int) (*[]entity.BeatmapSet, error)
	ComputeBeatmapSet(set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error)

	GetLatestBeatmapId() (uint, error)
	GetBeatmapSetIdForUpdate(limit int) ([]uint, error)
	CreateBeatmapSet(from interface{}) (*entity.BeatmapSetFull, error)
	UpdateBeatmapSet(id uint, from interface{}) (*entity.BeatmapSetFull, error)
	DeleteBeatmapSet(id uint) error

	Fetch(id uint) (*entity.BeatmapSetFull, error)
}
