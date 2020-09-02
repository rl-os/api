package sql

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/store"
	"time"
)

type BeatmapSetStore struct {
	SqlStore
}

func newSqlBeatmapSetStore(sqlStore SqlStore) store.BeatmapSet {
	return &BeatmapSetStore{sqlStore}
}

func (s BeatmapSetStore) SetFavourite(ctx context.Context, userId uint, id uint) (uint, error) {
	_, err := s.GetMaster().ExecContext(
		ctx,
		`insert into favouritemaps (beatmapset_id, user_id)
		select $1, $2
		where not exists (
			select id from favouritemaps where beatmapset_id = $1 and user_id = $2
		)`,
		id,
		userId,
	)
	if err != nil {
		return 0, errors.WithCause(
			"bms_set_favourite",
			400,
			"beatmapset not found or already favourited",
			err,
		)
	}

	var total uint
	err = s.GetMaster().GetContext(
		ctx,
		&total,
		`select count(id) as total from favouritemaps where user_id = $1`,
		userId,
	)
	if err != nil {
		return 0, errors.WithCause(
			"bms_set_favourite",
			404,
			"favourited maps not found",
			err,
		)
	}

	return total, nil
}

func (s BeatmapSetStore) SetUnFavourite(ctx context.Context, userId uint, id uint) (uint, error) {
	_, err := s.GetMaster().ExecContext(
		ctx,
		`delete from favouritemaps where beatmapset_id = $1 and user_id = $2`,
		id,
		userId,
	)
	if err != nil {
		return 0, errors.WithCause(
			"bms_set_favourite",
			400,
			"beatmapset not found or already favourited",
			err,
		)
	}

	var total uint
	err = s.GetMaster().GetContext(
		ctx,
		&total,
		`select count(id) as total from favouritemaps where user_id = $1`,
		userId,
	)
	if err != nil {
		return 0, errors.WithCause(
			"bms_set_favourite",
			404,
			"favourited maps not found",
			err,
		)
	}

	return total, nil
}

func (s BeatmapSetStore) Get(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	set := &entity.BeatmapSetFull{}

	userId, ok := ctx.Value("current_user_id").(uint)
	if !ok {
		userId = 0
	}

	err := s.GetMaster().GetContext(
		ctx,
		set,
		`SELECT id, last_checked, title, artist,
       		submitted_date, last_updated, ranked_date,
			creator, user_id, bpm, source, covers, preview_url, tags, video,
			storyboard, ranked, status, is_scoreable, discussion_enabled,
			discussion_locked, can_be_hyped, availability, hype, nominations,
			legacy_thread_url, description, genre, language, "user",
        	coalesce((
        	    select true from favouritemaps where beatmapset_id = $1 and user_id = $2
        	), false) as has_favourited,
       		coalesce((
       		    select count(*) from favouritemaps where beatmapset_id = $1
       		), 0) as favourite_count,
       		0 as play_count
		FROM beatmap_set
		WHERE id = $1;`,
		id,
		userId,
	)

	switch err {
	case sql.ErrNoRows:
		data, err := s.BeatmapSet().FetchFromBancho(ctx, id)
		if err != nil {
			return nil, err
		}

		set, err = s.BeatmapSet().Create(ctx, data)
		if err != nil {
			return nil, err
		}
	}

	return s.BeatmapSet().ComputeFields(ctx, *set)
}

func (s BeatmapSetStore) GetAll(ctx context.Context, page int, limit int) (*[]entity.BeatmapSet, error) {
	panic("implement me")
}

func (s BeatmapSetStore) Create(ctx context.Context, from interface{}) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	b, err := json.Marshal(&from)
	if err != nil {
		return nil, errors.WithCause("bms_create", 500, "marshaling input interface", err)
	}

	err = s.GetMaster().GetContext(
		ctx,
		&set,
		`insert into beatmap_set
		select id, last_checked, title, artist, play_count, favourite_count,
			submitted_date, last_updated, ranked_date,
			creator, user_id, bpm, source, covers, preview_url, tags, video,
			storyboard, ranked, status, is_scoreable, discussion_enabled,
			discussion_locked, can_be_hyped, availability, hype, nominations,
			legacy_thread_url, description, genre, language, "user"
		from json_populate_record(NULL::beatmap_set, $1)
		returning *`,
		string(b),
	)
	if err != nil {
		return nil, errors.WithCause("bms_create", 400, "beatmapset not created", err)
	}

	// todo: может отказаться от from interface{} ?
	// тк придется поддерживать все необходимые структуры
	if m, ok := from.(*entity.BeatmapSetFull); ok {
		data, err := s.Beatmap().CreateBatch(ctx, m.Beatmaps)
		if err == nil {
			set.Beatmaps = *data
		}
	}

	return &set, nil
}

func (s BeatmapSetStore) Update(ctx context.Context, id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	var set entity.BeatmapSetFull

	b, err := json.Marshal(&from)
	if err != nil {
		return nil, errors.WithCause("bms_update", 500, "marshaling input interface", err)
	}

	// update only required fields from json
	// easy to change
	err = s.GetMaster().GetContext(
		ctx,
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
		return nil, errors.WithCause("bms_update", 400, "beatmapset not updated", err)
	}

	return &set, nil
}

func (s BeatmapSetStore) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

// FetchFromBancho beatmapset from original api
func (s BeatmapSetStore) FetchFromBancho(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	data, err := s.GetOsuClient().BeatmapSet.Get(id)
	if err != nil {
		return nil, errors.WithCause("bms_fetch", 404, "beatmapset not found", err)
	}

	var out entity.BeatmapSetFull
	if err := copier.Copy(&out, &data); err != nil {
		return nil, errors.WithCause("bms_fetch", 500, "invalid structs", err)
	}

	return &out, nil
}

// GetIdsForUpdate and return list of ids
func (s BeatmapSetStore) GetIdsForUpdate(ctx context.Context, limit int) ([]uint, error) {
	ids := make([]uint, 0)
	now := time.Now()

	err := s.GetMaster().SelectContext(
		ctx,
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
	if err != nil {
		return nil, errors.WithCause("bms_need_update", 404, "beatmapset not found", err)
	}

	return ids, nil
}

func (s BeatmapSetStore) GetLatestId(ctx context.Context) (uint, error) {
	var id uint

	err := s.GetMaster().GetContext(ctx, &id, `select id from beatmap_set order by id desc`)

	switch err {
	case sql.ErrNoRows:
		return 0, nil
	default:
		return id, err
	}
}

func (s BeatmapSetStore) ComputeFields(ctx context.Context, set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error) {
	set.RecentFavourites = make([]entity.UserShortField, 0)
	set.Ratings = make([]int64, 11)
	set.Converts = []entity.Beatmap{}
	set.Beatmaps = s.Beatmap().GetBySetId(ctx, uint(set.ID))

	for _, b := range set.Beatmaps {
		if b.Mode != entity.Osu {
			continue
		}

		for i, mode := range entity.Modes {
			if mode == entity.Osu {
				continue
			}

			// todo: calculate difficulties
			conv := b
			conv.Convert = true
			conv.Mode = mode
			conv.ModeInt = int64(i)

			set.Converts = append(set.Converts, conv)
		}
	}

	return &set, nil
}
