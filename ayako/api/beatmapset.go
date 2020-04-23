package api

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BeatmapSetHandlers struct {
	Store store.Store
}

func (api *Routes) InitBeatmapSet(store store.Store) {
	handlers := BeatmapSetHandlers{store}

	api.BeatmapSets.GET("/:beatmapset", handlers.Get)
}

func (h *BeatmapSetHandlers) Get(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmapset id")
	}

	beatmaps, err := h.Store.BeatmapSet().GetBeatmapSet(uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	return c.JSON(200, beatmaps)
}
