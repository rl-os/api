package api

import (
	"github.com/deissh/osu-lazer/ayako/entity"
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

	api.Beatmaps.GET("/:id", handlers.Show)
	api.Beatmaps.GET("/lookup", handlers.Lookup)
}

func (h *BeatmapHandlers) Show(c echo.Context) error {
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

func (h *BeatmapHandlers) Lookup(c echo.Context) (err error) {
	params := struct {
		Id       uint   `query:"id"`
		CheckSum string `query:"checksum"`
		Filename string `query:"filename"`
	}{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	var beatmap *entity.SingleBeatmap
	if params.CheckSum != "" {
		// todo: search by md5
		beatmap, err = h.Store.Beatmap().GetBeatmap(params.Id)
	}

	if beatmap == nil && params.Id != 0 {
		beatmap, err = h.Store.Beatmap().GetBeatmap(params.Id)
	}

	if beatmap == nil && params.Filename != "" {
		// todo: search by filename
		beatmap, err = h.Store.Beatmap().GetBeatmap(params.Id)
	}

	if err != nil || beatmap == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	return c.JSON(200, beatmap)
}
