package sql

import (
	"encoding/json"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"time"
)

type BeatmapSetStore struct {
	SqlStore
}

func newSqlBeatmapSetStore(sqlStore SqlStore) store.BeatmapSet {
	return &BeatmapSetStore{sqlStore}
}

func (s BeatmapSetStore) GetBeatmapSet(id uint) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	err := s.GetMaster().Get(
		&set,
		`SELECT id, title, artist, play_count, favourite_count,
			has_favourited, submitted_date, last_updated, ranked_date,
		   creator, user_id, bpm, source, covers, preview_url, tags, video,
		   storyboard, ranked, status, is_scoreable, discussion_enabled,
		   discussion_locked, can_be_hyped, availability, hype, nominations,
		   legacy_thread_url, description, genre, language, "user"
		FROM beatmap_set
		WHERE id = $1;`,
		id,
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("store.GetBeatmapSet")

		//todo: error wrap
		return nil, err
	}

	return &set, nil
}

func (s BeatmapSetStore) GetAllBeatmapSets(page int, limit int) (*[]entity.BeatmapSet, error) {
	panic("implement me")
}

func (s BeatmapSetStore) CreateBeatmapSet(from interface{}) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	b, err := json.Marshal(&from)
	if err != nil {
		log.Error().
			Err(err).
			Msg("store.CreateBeatmapSet")

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
		log.Error().
			Err(err).
			Msg("store.CreateBeatmapSet")

		//todo: error wrap
		return nil, err
	}

	return &set, nil
}

func (s BeatmapSetStore) UpdateBeatmapSet(id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	b, err := json.Marshal(&from)
	if err != nil {
		log.Error().
			Err(err).
			Msg("store.CreateBeatmapSet")

		return nil, err
	}

	// update only required files from json
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
		log.Error().
			Err(err).
			Msg("store.UpdateBeatmapSet")

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
		log.Error().
			Err(err).
			Int64("id", data.ID).
			Msg("scan to entity.BeatmapSetFull")
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
		now.Truncate(time.Hour*24*7),
		now.Truncate(time.Hour),
		limit,
	)
	return ids, err
}
