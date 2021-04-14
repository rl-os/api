package gorm

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"gorm.io/gorm/clause"
)

type BeatmapRepository struct {
	*Supplier
}

func NewBeatmapRepository(supplier *Supplier) repository.Beatmap {
	return &BeatmapRepository{supplier}
}

func (s BeatmapRepository) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
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

func (s BeatmapRepository) GetBySetId(ctx context.Context, beatmapsetId uint) (*[]entity.Beatmap, error) {
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

func (s BeatmapRepository) Create(ctx context.Context, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapRepository) CreateBatch(ctx context.Context, from interface{}) (*[]entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapRepository) Update(ctx context.Context, id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapRepository) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
