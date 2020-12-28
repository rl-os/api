package sql

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"gorm.io/gorm/clause"
)

type BeatmapSetStore struct {
	SqlStore
}

func newSqlBeatmapSetStore(sqlStore SqlStore) store.BeatmapSet {
	return &BeatmapSetStore{sqlStore}
}

func (s BeatmapSetStore) Update(ctx context.Context, id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s BeatmapSetStore) SetFavourite(ctx context.Context, userId uint, id uint) (uint, error) {
	panic("implement me")
}

func (s BeatmapSetStore) SetUnFavourite(ctx context.Context, userId uint, id uint) (uint, error) {
	panic("implement me")
}

func (s BeatmapSetStore) Get(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	bms := entity.BeatmapSetFull{}

	err := s.GetMaster().
		WithContext(ctx).
		Table("beatmap_set").
		Where("id = ?", id).
		Preload(clause.Associations).
		First(&bms).
		Error

	if err != nil {
		return nil, err
	}

	return &bms, nil
}

func (s BeatmapSetStore) GetAll(ctx context.Context, page int, limit int) (*[]entity.BeatmapSet, error) {
	panic("implement me")
}

func (s BeatmapSetStore) Create(ctx context.Context, from interface{}) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s BeatmapSetStore) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

// FetchFromBancho beatmapset from original api
func (s BeatmapSetStore) FetchFromBancho(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

// GetIdsForUpdate and return list of ids
func (s BeatmapSetStore) GetIdsForUpdate(ctx context.Context, limit int) ([]uint, error) {
	panic("implement me")
}

func (s BeatmapSetStore) GetLatestId(ctx context.Context) (uint, error) {
	panic("implement me")
}

func (s BeatmapSetStore) ComputeFields(ctx context.Context, set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}
