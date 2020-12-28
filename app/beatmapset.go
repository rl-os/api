package app

import (
	"context"
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

// GetBeatmapset from store and return 404 error if not exist
func (a *App) GetBeatmapset(ctx context.Context, beatmapsetID uint) (*entity.BeatmapSetFull, error) {
	data, err := a.Store.BeatmapSet().Get(ctx, beatmapsetID)
	if err != nil {
		return nil, ErrNotFoundBMS.WithCause(err)
	}

	return data, nil
}

func (a *App) LookupBeatmapset(ctx context.Context, beatmapId uint) (*entity.BeatmapSetFull, error) {
	beatmap, err := a.Store.Beatmap().Get(ctx, beatmapId)
	if err != nil {
		return nil, err
	}

	return a.Store.BeatmapSet().Get(ctx, uint(beatmap.Beatmapset.ID))
}

func (a *App) SearchBeatmapset(ctx context.Context) (*entity.BeatmapsetSearchResult, error) {
	// TODO: this
	return nil, nil
}

func (a *App) FavouriteBeatmapset(ctx context.Context, action string, beatmapsetID uint) (uint, error) {
	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return 0, ErrNotFoundBMS.WithCause(err)
	}

	var count uint
	store := a.Store.BeatmapSet()

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
