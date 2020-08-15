package app

import (
	"context"
	myctx "github.com/deissh/rl/ayako/ctx"
	"github.com/deissh/rl/ayako/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *App) GetBeatmapset(ctx context.Context, beatmapsetID uint) (*entity.BeatmapSetFull, error) {
	// todo: check error
	return a.Store.BeatmapSet().Get(ctx, beatmapsetID)
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
		return 0, err
	}

	switch action {
	case "favourite":
		return a.Store.BeatmapSet().SetFavourite(ctx, userId, beatmapsetID)
	case "unfavourite":
		return a.Store.BeatmapSet().SetFavourite(ctx, userId, beatmapsetID)
	default:
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
	}
}
