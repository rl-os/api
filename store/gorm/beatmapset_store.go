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

func (s BeatmapSetStore) SetFavourite(ctx context.Context, userId uint, id uint) (uint, error) {
	_ = s.GetMaster().WithContext(ctx).
		Table("user_beatmapset_favourite").
		Create(map[string]interface{}{
			"beatmapset_id": id,
			"user_id":       userId,
		})

	var count int64
	err := s.GetMaster().
		Table("user_beatmapset_favourite").
		Where("beatmapset_id = ? AND user_id = ?", id, userId).
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}

	return uint(count), nil
}

func (s BeatmapSetStore) SetUnFavourite(ctx context.Context, userId uint, id uint) (uint, error) {
	err := s.GetMaster().
		Exec("DELETE FROM user_beatmapset_favourite WHERE beatmapset_id = ? AND user_id = ?", id, userId).
		Error
	if err != nil {
		return 0, err
	}

	var count int64
	err = s.GetMaster().
		Table("user_beatmapset_favourite").
		Where("beatmapset_id = ? AND user_id = ?", id, userId).
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}

	return uint(count), nil
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

func (s BeatmapSetStore) Create(ctx context.Context, from interface{}) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s BeatmapSetStore) Update(ctx context.Context, id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	panic("implement me")
}

func (s BeatmapSetStore) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
