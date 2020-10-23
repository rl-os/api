package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
)

type BeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &BeatmapStore{sqlStore}
}

func (s BeatmapStore) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) GetBySetId(ctx context.Context, beatmapsetId uint) []entity.Beatmap {
	panic("implement me")
}

func (s BeatmapStore) Create(ctx context.Context, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) CreateBatch(ctx context.Context, from interface{}) (*[]entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) Update(ctx context.Context, id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
