package sql

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
)

type BeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &BeatmapStore{sqlStore}
}

func (s BeatmapStore) GetBeatmap(id uint) (*entity.Beatmap, error) {
	var beatmap entity.Beatmap

	return &beatmap, nil
}

func (s BeatmapStore) GetAllBeatmap(page int, limit int) (*[]entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) CreateBeatmap(from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) UpdateBeatmap(id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) DeleteBeatmap(id uint) error {
	panic("implement me")
}
