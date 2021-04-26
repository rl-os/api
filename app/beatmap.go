package app

import (
	"context"
	"github.com/rl-os/api/repository"
	"net/http"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
)

var (
	ErrNotFoundBM = errors.New("beatmap", http.StatusNotFound, "Beatmap not found")
)

type BeatmapUseCase struct {
	*App
	BeatmapRepository repository.Beatmap
}

func NewBeatmapUseCase(app *App, rep repository.Beatmap) *BeatmapUseCase {
	return &BeatmapUseCase{app, rep}
}

// GetBeatmap from repository
func (a *BeatmapUseCase) GetBeatmap(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	data, err := a.BeatmapRepository.Get(ctx, id)
	if err != nil {
		return nil, ErrNotFoundBM.WithCause(err)
	}

	return data, nil
}

func (a *BeatmapUseCase) LookupBeatmap(ctx context.Context, id uint, checksum string, filename string) (*entity.SingleBeatmap, error) {
	if checksum != "" {
		// todo: search by md5
		return nil, ErrNotFoundBM
	}

	if id != 0 {
		beatmap, err := a.BeatmapRepository.Get(ctx, id)
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
