package sql

import "github.com/deissh/osu-lazer/ayako/store"

type SqlBeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &SqlBeatmapStore{sqlStore}
}

func (s SqlBeatmapStore) GetBeatmap(id uint) interface{} {
	panic("implement me")
}

func (s SqlBeatmapStore) GetAllBeatmap(page int, limit int) []interface{} {
	panic("implement me")
}

func (s SqlBeatmapStore) CreateBeatmap(from interface{}) interface{} {
	panic("implement me")
}

func (s SqlBeatmapStore) UpdateBeatmap(id uint, from interface{}) interface{} {
	panic("implement me")
}

func (s SqlBeatmapStore) DeleteBeatmap(id uint) {
	panic("implement me")
}
