package client

import (
	"errors"
	"fmt"
	"github.com/deissh/osu-lazer/ayako/entity"
)

type BeatmapSetAPI struct {
	*OsuAPI
}

func (b *BeatmapSetAPI) Get(id uint) (*entity.BeatmapSetFull, error) {
	json := entity.BeatmapSetFull{}

	req := b.client.
		Request().
		Path(fmt.Sprint("/api/v2/beatmapsets/", id)).
		Use(b.bearerMiddleware())

	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.RawResponse.Status)
	}

	if err := res.JSON(&json); err != nil {
		return nil, err
	}

	return &json, nil
}
