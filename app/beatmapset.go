package app

import (
	"context"
	"github.com/rl-os/api/repository"
	"net/http"
	"strconv"

	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

// BMS = beatmapset
var (
	ErrNotFoundBMS      = errors.New("beatmapset", http.StatusNotFound, "Beatmapset not found")
	ErrInvalidBMSAction = errors.New("beatmapset", http.StatusBadRequest, "Invalid action")
)

type BeatmapSetUseCase struct {
	*App
	beatmapUseCase *BeatmapUseCase
	repository     repository.BeatmapSet
}

func NewBeatmapSetUseCase(
	app *App,
	beatmapUseCase *BeatmapUseCase,
	repository repository.BeatmapSet,
) *BeatmapSetUseCase {
	return &BeatmapSetUseCase{app, beatmapUseCase, repository}
}

// Get from repository and return 404 error if not exist
func (a *BeatmapSetUseCase) Get(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	var cached entity.BeatmapSetFull

	err := a.Cache.Get("beatmapset", strconv.Itoa(int(id)), &cached)
	if err == nil {
		return &cached, nil
	}

	data, err := a.repository.Get(ctx, id)
	if err != nil {
		return nil, ErrNotFoundBMS.WithCause(err)
	}

	a.Cache.Set("beatmapset", strconv.Itoa(int(id)), data)

	return data, nil
}

func (a *BeatmapSetUseCase) Lookup(ctx context.Context, beatmapId uint) (*entity.BeatmapSetFull, error) {
	beatmap, err := a.beatmapUseCase.Get(ctx, beatmapId)
	if err != nil {
		return nil, ErrNotFoundBM.WithCause(err)
	}

	var cached entity.BeatmapSetFull

	err = a.Cache.Get("beatmapset", strconv.Itoa(int(beatmap.Beatmapset.ID)), &cached)
	if err == nil {
		return &cached, nil
	}

	set, err := a.repository.Get(ctx, uint(beatmap.Beatmapset.ID))
	if err != nil {
		return nil, ErrNotFoundBMS.WithCause(err)
	}

	a.Cache.Set("beatmapset", strconv.Itoa(int(set.ID)), set)

	return set, nil
}

func (a *BeatmapSetUseCase) Search(ctx context.Context) (*entity.BeatmapsetSearchResult, error) {
	// TODO: this
	return nil, nil
}

func (a *BeatmapSetUseCase) Favourite(ctx context.Context, action string, beatmapsetID uint) (uint, error) {
	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return 0, ErrNotFoundBMS.WithCause(err)
	}

	var count uint
	store := a.repository

	switch action {
	case "favourite":
		count, err = store.SetFavourite(ctx, userId, beatmapsetID)
	case "unfavourite":
		count, err = store.SetUnFavourite(ctx, userId, beatmapsetID)
	default:
		return 0, ErrInvalidBMSAction
	}

	return count, nil
}
