package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"gorm.io/gorm/clause"
)

type BeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &BeatmapStore{sqlStore}
}

func (s BeatmapStore) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	bm := entity.SingleBeatmap{}

	err := s.GetMaster().
		WithContext(ctx).
		Table("beatmaps").
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&bm).
		Error

	if err != nil {
		return nil, err
	}

	return &bm, nil
}

func (s BeatmapStore) GetBySetId(ctx context.Context, beatmapsetId uint) (*[]entity.Beatmap, error) {
	var bms []entity.Beatmap

	err := s.GetMaster().
		WithContext(ctx).
		Table("beatmaps").
		Where("beatmapset_id = ?", beatmapsetId).
		Preload(clause.Associations).
		Find(&bms).
		Error

	if err != nil {
		return nil, err
	}

	return &bms, nil
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
