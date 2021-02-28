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

type Beatmap struct {
	*App
}

// Get from store
func (b *Beatmap) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	data, err := b.Store.Beatmap().Get(ctx, id)
	if err != nil {
		return nil, ErrNotFoundBM.WithCause(err)
	}

	return data, nil
}

func (b *Beatmap) Lookup(ctx context.Context, id uint, checksum string, filename string) (*entity.SingleBeatmap, error) {
	if checksum != "" {
		// todo: search by md5
		return nil, ErrNotFoundBM
	}

	if id != 0 {
		beatmap, err := b.Store.Beatmap().Get(ctx, id)
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
