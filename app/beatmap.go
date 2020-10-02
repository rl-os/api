package app

import (
	"context"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

// BM = beatmap
var (
	ErrNotFoundBM = errors.New("beatmap", http.StatusNotFound, "Beatmap not found")
)

// GetBeatmap from store
func (a *App) GetBeatmap(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	data, err := a.Store.Beatmap().Get(ctx, id)
	if err != nil {
		return nil, ErrNotFoundBM.WithCause(err)
	}

	return data, nil
}

func (a *App) LookupBeatmap(ctx context.Context, id uint, checksum string, filename string) (*entity.SingleBeatmap, error) {
	if checksum != "" {
		// todo: search by md5
		return nil, ErrNotFoundBM
	}

	if id != 0 {
		beatmap, err := a.Store.Beatmap().Get(ctx, id)
		if err != nil {
			return nil, ErrNotFoundBM
		}

		return beatmap, nil
	}

	if filename != "" {
		// todo: search by filename
		return nil, ErrNotFoundBM
	}

	return nil, ErrNotFoundBM
}
