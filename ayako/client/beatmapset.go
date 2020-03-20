package client

import (
	"errors"
	"fmt"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/rs/zerolog/log"
)

type BeatmapSetAPI struct {
	*OsuAPI
}

func (b *BeatmapSetAPI) Get(id uint) (*entity.BeatmapSetFull, error) {
	json := entity.BeatmapSetFull{}

	req := b.Client.
		Request().
		Path(fmt.Sprint("/api/v2/beatmapsets/", id))

	res, err := req.Send()
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	if !res.Ok {
		log.Error().
			Str("status", res.RawResponse.Status).
			Msg("Request not ok")
		return nil, errors.New(res.RawResponse.Status)
	}

	if err := res.JSON(&json); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &json, nil
}
