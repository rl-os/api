package api

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
)

type BeatmapHandlers struct {
	Store store.Store
}

func (api *Routes) InitBeatmaps(store store.Store) {
	handlers := BeatmapHandlers{store}

	api.Beatmaps.GET("", handlers.LookUp)
}

func (h *BeatmapHandlers) LookUp(c echo.Context) error {
	beatmaps, _ := h.Store.Beatmap().GetBeatmap(1)

	return c.JSON(200, beatmaps)
}
