package sql

import (
	"context"
	"encoding/json"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/store"
)

type BeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &BeatmapStore{sqlStore}
}

func (s BeatmapStore) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	var beatmap entity.SingleBeatmap

	err := s.GetMaster().GetContext(
		ctx,
		&beatmap,
		`select id, beatmapset_id, mode, mode_int, convert,
       	  difficulty_rating, version, total_length, hit_length,
          bpm, cs, drain, accuracy, ar, playcount, passcount,
          count_circles, count_sliders, count_spinners, count_total,
          is_scoreable, last_updated, ranked, status, url, deleted_at, max_combo
		from beatmaps
		where id = $1`,
		id,
	)
	if err != nil {
		return nil, errors.WithCause("bm_get", 404, "Beatmap not found", err)
	}
	set, err := s.BeatmapSet().Get(ctx, uint(beatmap.BeatmapsetID))
	if err != nil {
		return nil, err
	}

	beatmap.Beatmapset = set.BeatmapSet

	return &beatmap, nil
}

func (s BeatmapStore) GetBySetId(ctx context.Context, beatmapsetId uint) []entity.Beatmap {
	beatmaps := make([]entity.Beatmap, 0)

	_ = s.GetMaster().SelectContext(
		ctx,
		&beatmaps,
		`select id, beatmapset_id, mode, mode_int, convert,
		   difficulty_rating, version, total_length,
		   hit_length, bpm, cs, drain, accuracy, ar,
		   playcount, passcount, count_circles,
		   count_sliders, count_spinners, count_total,
		   is_scoreable, last_updated, ranked,
		   status, url, deleted_at, max_combo
		from beatmaps
		where beatmapset_id = $1`,
		beatmapsetId,
	)

	return beatmaps
}

func (s BeatmapStore) Create(ctx context.Context, from interface{}) (*entity.Beatmap, error) {
	var set entity.Beatmap

	b, err := json.Marshal(&from)
	if err != nil {
		return nil, errors.WithCause("bm_create", 500, "marshaling input interface", err)
	}

	err = s.GetMaster().GetContext(
		ctx,
		&set,
		`insert into beatmaps
		select id, beatmapset_id, mode, mode_int, convert,
		       difficulty_rating, version, total_length,
		       hit_length, bpm, cs, drain, accuracy, ar,
		       playcount, passcount, count_circles,
		       count_sliders, count_spinners, count_total,
		       is_scoreable, last_updated, ranked,
		       status, url, deleted_at, max_combo
		from json_populate_record(NULL::beatmaps, $1)
		returning *`,
		string(b),
	)
	if err != nil {
		return nil, errors.WithCause("bm_create", 400, "beatmap not created", err)
	}

	return &set, nil
}

func (s BeatmapStore) CreateBatch(ctx context.Context, from interface{}) (*[]entity.Beatmap, error) {
	var sets []entity.Beatmap

	b, err := json.Marshal(&from)
	if err != nil {
		return nil, errors.WithCause("bm_create", 500, "marshaling input interface", err)
	}

	err = s.GetMaster().SelectContext(
		ctx,
		&sets,
		`insert into beatmaps
		select id, beatmapset_id, mode, mode_int, convert,
		       difficulty_rating, version, total_length,
		       hit_length, bpm, cs, drain, accuracy, ar,
		       playcount, passcount, count_circles,
		       count_sliders, count_spinners, count_total,
		       is_scoreable, last_updated, ranked,
		       status, url, deleted_at, max_combo
		from json_populate_recordset(NULL::beatmaps, $1)
		returning *`,
		string(b),
	)
	if err != nil {
		return nil, errors.WithCause("bm_create_batch", 400, "beatmap not created", err)
	}

	return &sets, nil
}

func (s BeatmapStore) Update(ctx context.Context, id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
