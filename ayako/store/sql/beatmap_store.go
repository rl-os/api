package sql

import (
	"encoding/json"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
)

type BeatmapStore struct {
	SqlStore
}

func newSqlBeatmapStore(sqlStore SqlStore) store.Beatmap {
	return &BeatmapStore{sqlStore}
}

func (s BeatmapStore) Get(id uint) (*entity.SingleBeatmap, error) {
	var beatmap entity.SingleBeatmap

	err := s.GetMaster().Get(
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
		return nil, err
	}
	set, err := s.BeatmapSet().Get(uint(beatmap.BeatmapsetID))
	if err != nil {
		return nil, err
	}

	beatmap.Beatmapset = set.BeatmapSet

	return &beatmap, nil
}

func (s BeatmapStore) GetBySetId(beatmapsetId uint) []entity.Beatmap {
	beatmaps := make([]entity.Beatmap, 0)

	_ = s.GetMaster().Select(
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

func (s BeatmapStore) Create(from interface{}) (*entity.Beatmap, error) {
	var set entity.Beatmap

	b, err := json.Marshal(&from)
	if err != nil {
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
		return nil, err
	}

	return &set, nil
}

func (s BeatmapStore) CreateBatch(from interface{}) (*[]entity.Beatmap, error) {
	var sets []entity.Beatmap

	b, err := json.Marshal(&from)
	if err != nil {
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
		return nil, err
	}

	return &sets, nil
}

func (s BeatmapStore) Update(id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (s BeatmapStore) Delete(id uint) error {
	panic("implement me")
}
