package api

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
)

type BeatmapSetHandlers struct {
	Store store.Store
}

func (api *Routes) InitBeatmapSet(store store.Store) {
	handlers := BeatmapSetHandlers{store}

	api.BeatmapSets.GET("", handlers.Get)
}

func (h *BeatmapSetHandlers) Get(c echo.Context) error {
	beatmaps, _ := h.Store.BeatmapSet().GetBeatmapSet(23416)

	return c.JSON(200, beatmaps)
}
