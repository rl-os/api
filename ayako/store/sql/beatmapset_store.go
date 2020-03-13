package sql

import "github.com/deissh/osu-lazer/ayako/store"

type SqlBeatmapSetStore struct {
	SqlStore
}

func newSqlBeatmapSetStore(sqlStore SqlStore) store.BeatmapSet {
	return &SqlBeatmapSetStore{sqlStore}
}

func (s SqlBeatmapSetStore) Get(id uint) interface{} {
	panic("implement me")
}

func (s SqlBeatmapSetStore) GetAll(page int, limit int) []interface{} {
	panic("implement me")
}

func (s SqlBeatmapSetStore) Create(from interface{}) interface{} {
	panic("implement me")
}

func (s SqlBeatmapSetStore) Update(id uint, from interface{}) interface{} {
	panic("implement me")
}

func (s SqlBeatmapSetStore) Delete(id uint) {
	panic("implement me")
}
