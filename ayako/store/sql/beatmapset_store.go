package sql

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
)

type SqlBeatmapSetStore struct {
	SqlStore
}

func newSqlBeatmapSetStore(sqlStore SqlStore) store.BeatmapSet {
	return &SqlBeatmapSetStore{sqlStore}
}

func (s SqlBeatmapSetStore) GetBeatmapSet(id uint) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s SqlBeatmapSetStore) GetAllBeatmapSets(page int, limit int) (*[]entity.BeatmapSet, error) {
	panic("implement me")
}

func (s SqlBeatmapSetStore) CreateBeatmapSet(from interface{}) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s SqlBeatmapSetStore) UpdateBeatmapSet(id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s SqlBeatmapSetStore) DeleteBeatmapSet(id uint) error {
	panic("implement me")
}
