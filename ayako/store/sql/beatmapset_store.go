package sql

import (
	"database/sql"
	"encoding/json"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/jinzhu/copier"
	"time"
)

type BeatmapSetStore struct {
	SqlStore
}

func newSqlBeatmapSetStore(sqlStore SqlStore) store.BeatmapSet {
	return &BeatmapSetStore{sqlStore}
}

func (s BeatmapSetStore) GetBeatmapSet(id uint) (*entity.BeatmapSetFull, error) {
	set := &entity.BeatmapSetFull{}

	err := s.GetMaster().Get(
		set,
		`SELECT id, last_checked, title, artist, play_count, favourite_count,
			has_favourited, submitted_date, last_updated, ranked_date,
		   creator, user_id, bpm, source, covers, preview_url, tags, video,
		   storyboard, ranked, status, is_scoreable, discussion_enabled,
		   discussion_locked, can_be_hyped, availability, hype, nominations,
		   legacy_thread_url, description, genre, language, "user"
		FROM beatmap_set
		WHERE id = $1;`,
		id,
	)

	switch err {
	case sql.ErrNoRows:
		data, err := s.Fetch(id)
		if err != nil {
			return nil, err
		}

		set, err = s.CreateBeatmapSet(data)
		if err != nil {
			return nil, err
		}
	}

	return s.ComputeBeatmapSet(*set)
}

func (s BeatmapSetStore) ComputeBeatmapSet(set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error) {
	set.RecentFavourites = []entity.User{}
	set.Ratings = make([]int64, 11)
	set.Converts = []entity.Beatmap{}
	set.Beatmaps = s.Beatmap().GetBeatmapsBySet(uint(set.ID))

	return &set, nil
}

func (s BeatmapSetStore) GetAllBeatmapSets(page int, limit int) (*[]entity.BeatmapSet, error) {
	panic("implement me")
}

func (s BeatmapSetStore) CreateBeatmapSet(from interface{}) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	b, err := json.Marshal(&from)
	if err != nil {
		return nil, err
	}

	err = s.GetMaster().Get(
		&set,
		`insert into beatmap_set
		select id, last_checked, title, artist, play_count, favourite_count,
			has_favourited, submitted_date, last_updated, ranked_date,
		   creator, user_id, bpm, source, covers, preview_url, tags, video,
		   storyboard, ranked, status, is_scoreable, discussion_enabled,
		   discussion_locked, can_be_hyped, availability, hype, nominations,
		   legacy_thread_url, description, genre, language, "user"
		from json_populate_record(NULL::beatmap_set, $1)
		returning *`,
		string(b),
	)
	if err != nil {
		return nil, err
	}

	// todo: может отказаться от from interface{} ?
	// тк придется поддерживать все необходимые структуры
	if m, ok := from.(*entity.BeatmapSetFull); ok {
		data, err := s.Beatmap().CreateBeatmaps(m.Beatmaps)
		if err == nil {
			set.Beatmaps = *data
		}
	}

	return &set, nil
}

func (s BeatmapSetStore) UpdateBeatmapSet(id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	b, err := json.Marshal(&from)
	if err != nil {
		return nil, err
	}

	// update only required fields from json
	// easy to change
	err = s.GetMaster().Get(
		&set,
		`update beatmap_set set
			last_checked = sq.last_checked, last_updated = sq.last_updated,
            title = sq.title, artist = sq.artist, submitted_date = sq.submitted_date,
            ranked_date = sq.ranked_date, bpm = sq.bpm, source = sq.source, covers = sq.covers,
            preview_url = sq.preview_url, tags = sq.tags, video = sq.video,
            storyboard = sq.storyboard, ranked = sq.ranked, status = sq.status,
            description = sq.description, genre = sq.genre, language = sq.language
		from (
		    select last_checked, last_updated, title, artist, submitted_date, ranked_date,
		           bpm, source, covers, preview_url, tags, video,
		           storyboard, ranked, status, description, genre, language
		    from json_populate_record(null::beatmap_set, $2)
		) as sq
		where id = $1
		returning *`,
		id,
		string(b),
	)
	if err != nil {
		return nil, err
	}

	return &set, nil
}

func (s BeatmapSetStore) DeleteBeatmapSet(id uint) error {
	panic("implement me")
}

// Fetch beatmapset from original api
func (s BeatmapSetStore) Fetch(id uint) (*entity.BeatmapSetFull, error) {
	data, err := s.GetOsuClient().BeatmapSet.Get(id)
	if err != nil {
		return nil, err
	}

	var out entity.BeatmapSetFull
	if err := copier.Copy(&out, &data); err != nil {
		return nil, err
	}

	return &out, nil
}

// GetBeatmapSetIdForUpdate and return list of ids
func (s BeatmapSetStore) GetBeatmapSetIdForUpdate(limit int) ([]uint, error) {
	ids := make([]uint, 0)
	now := time.Now()

	err := s.GetMaster().Select(
		&ids,
		`select id
		from beatmap_set
		where (last_checked <= $1) or (last_checked <= $2 and status in ('pending', 'wip', 'qualified'))
		order by last_checked
		limit $3`,
		now.Add(-time.Hour*24*7),
		now.Add(-time.Minute*30),
		limit,
	)
	return ids, err
}

func (s BeatmapSetStore) GetLatestBeatmapId() (uint, error) {
	var id uint

	err := s.GetMaster().Get(&id, `select id from beatmap_set order by id desc`)

	switch err {
	case sql.ErrNoRows:
		return 0, nil
	default:
		return id, err
	}
}
