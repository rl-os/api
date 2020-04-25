package api

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BeatmapHandlers struct {
	Store store.Store
}

func (api *Routes) InitBeatmaps(store store.Store) {
	handlers := BeatmapHandlers{store}

	api.Beatmaps.GET("/:beatmap", handlers.Get)
}

func (h *BeatmapHandlers) Get(c echo.Context) error {
	beatmapID, err := strconv.ParseUint(c.Param("beatmap"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	beatmaps, err := h.Store.Beatmap().GetBeatmap(uint(beatmapID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	return c.JSON(200, beatmaps)
}
