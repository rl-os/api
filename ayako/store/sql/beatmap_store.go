package sql

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
)

type SqlBeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &SqlBeatmapStore{sqlStore}
}

func (s SqlBeatmapStore) GetBeatmap(id uint) (*entity.Beatmap, error) {
	var beatmap entity.Beatmap

	return &beatmap, nil
}

func (s SqlBeatmapStore) GetAllBeatmap(page int, limit int) (*[]entity.Beatmap, error) {
	panic("implement me")
}

func (s SqlBeatmapStore) CreateBeatmap(from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s SqlBeatmapStore) UpdateBeatmap(id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s SqlBeatmapStore) DeleteBeatmap(id uint) error {
	panic("implement me")
}
