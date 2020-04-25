package sql

import (
	"encoding/json"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/rs/zerolog/log"
)

type BeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &BeatmapStore{sqlStore}
}

func (s BeatmapStore) GetBeatmap(id uint) (*entity.Beatmap, error) {
	var beatmap entity.Beatmap

	return &beatmap, nil
}

func (s BeatmapStore) GetBeatmapsBySet(beatmapsetId uint) []entity.Beatmap {
	beatmaps := make([]entity.Beatmap, 0)

	err := s.GetMaster().Select(
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
	if err != nil {
		log.Error().
			Err(err).
			Msg("store.GetBeatmapsBySet")
	}

	return beatmaps
}

func (s BeatmapStore) GetAllBeatmap(page int, limit int) (*[]entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) CreateBeatmap(from interface{}) (*entity.Beatmap, error) {
	var set entity.Beatmap

	b, err := json.Marshal(&from)
	if err != nil {
		log.Error().
			Err(err).
			Msg("store.CreateBeatmap")

		return nil, err
	}

	err = s.GetMaster().Get(
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
		log.Error().
			Err(err).
			Msg("store.CreateBeatmap")

		return nil, err
	}

	return &set, nil
}

func (s BeatmapStore) CreateBeatmaps(from interface{}) (*[]entity.Beatmap, error) {
	var sets []entity.Beatmap

	b, err := json.Marshal(&from)
	if err != nil {
		log.Error().
			Err(err).
			Msg("store.CreateBeatmap")

		return nil, err
	}

	err = s.GetMaster().Select(
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
		log.Error().
			Err(err).
			Msg("store.CreateBeatmaps")

		return nil, err
	}

	return &sets, nil
}

func (s BeatmapStore) UpdateBeatmap(id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) DeleteBeatmap(id uint) error {
	panic("implement me")
}
