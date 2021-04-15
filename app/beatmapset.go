package app

import (
	"context"
	"github.com/rl-os/api/repository"
	"net/http"

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
	BeatmapRepository    repository.Beatmap
	BeatmapSetRepository repository.BeatmapSet
}

func NewBeatmapSetUseCase(app *App, beatmap repository.Beatmap, beatmapset repository.BeatmapSet) *BeatmapSetUseCase {
	return &BeatmapSetUseCase{app, beatmap, beatmapset}
}

// GetBeatmapset from repository and return 404 error if not exist
func (a *BeatmapSetUseCase) GetBeatmapset(ctx context.Context, beatmapsetID uint) (*entity.BeatmapSetFull, error) {
	data, err := a.BeatmapSetRepository.Get(ctx, beatmapsetID)
	if err != nil {
		return nil, ErrNotFoundBMS.WithCause(err)
	}

	return data, nil
}

func (a *BeatmapSetUseCase) LookupBeatmapset(ctx context.Context, beatmapId uint) (*entity.BeatmapSetFull, error) {
	beatmap, err := a.BeatmapRepository.Get(ctx, beatmapId)
	if err != nil {
		return nil, err
	}

	return a.BeatmapSetRepository.Get(ctx, uint(beatmap.Beatmapset.ID))
}

func (a *BeatmapSetUseCase) SearchBeatmapset(ctx context.Context) (*entity.BeatmapsetSearchResult, error) {
	// TODO: this
	return nil, nil
}

func (a *BeatmapSetUseCase) FavouriteBeatmapset(ctx context.Context, action string, beatmapsetID uint) (uint, error) {
	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return 0, ErrNotFoundBMS.WithCause(err)
	}

	var count uint
	store := a.BeatmapSetRepository

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
