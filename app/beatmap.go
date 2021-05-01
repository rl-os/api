package app

import (
	"context"
	"github.com/rl-os/api/repository"
	"net/http"
	"strconv"

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

// Get from repository
func (a *BeatmapUseCase) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	var cached entity.SingleBeatmap

	err := a.Cache.Get("beatmap", strconv.Itoa(int(id)), &cached)
	if err == nil {
		return &cached, nil
	}

	data, err := a.BeatmapRepository.Get(ctx, id)
	if err != nil {
		return nil, ErrNotFoundBM.WithCause(err)
	}

	a.Cache.Set("beatmap", strconv.Itoa(int(id)), data)

	return data, nil
}

func (a *BeatmapUseCase) Lookup(ctx context.Context, id uint, checksum string, filename string) (*entity.SingleBeatmap, error) {
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
