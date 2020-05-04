package api

import (
	"github.com/deissh/osu-lazer/ayako/entity"
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
	api.BeatmapSets.POST("/:beatmapset/favourites", handlers.Favourite)
	api.BeatmapSets.GET("/:beatmapset/download", handlers.Get)
	api.BeatmapSets.GET("/lookup", handlers.Lookup)
	api.BeatmapSets.GET("/search", handlers.Search)
}

func (h *BeatmapSetHandlers) Get(c echo.Context) error {
	beatmapsetID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmapset id")
	}

	beatmaps, err := h.Store.BeatmapSet().Get(uint(beatmapsetID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmapset not found")
	}

	return c.JSON(200, beatmaps)
}

func (h *BeatmapSetHandlers) Lookup(c echo.Context) (err error) {
	params := struct {
		Id uint `query:"beatmap_id"`
	}{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	beatmap, err := h.Store.Beatmap().Get(params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	beatmapSet, err := h.Store.BeatmapSet().Get(uint(beatmap.Beatmapset.ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "BeatmapSet not found")
	}

	return c.JSON(200, beatmapSet)
}

func (h *BeatmapSetHandlers) Search(c echo.Context) (err error) {
	// todo: this
	params := struct {
	}{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	beatmapSets, err := h.Store.BeatmapSet().Get(1118896)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "BeatmapSet not found")
	}

	return c.JSON(200, struct {
		Beatmapsets           *[]entity.BeatmapSetFull `json:"beatmapsets"`
		RecommendedDifficulty float32                  `json:"recommended_difficulty"`
		Error                 error                    `json:"error"`
		Total                 uint                     `json:"total"`
	}{
		&[]entity.BeatmapSetFull{*beatmapSets},
		3,
		nil,
		0,
	})
}

func (h *BeatmapSetHandlers) Favourite(c echo.Context) (err error) {
	params := struct {
		Action string `query:"action" json:"action"`
	}{}
	beatmapsetID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmapset id")
	}

	if err := c.Bind(&params); err != nil {
		return err
	}

	var total uint
	switch params.Action {
	case "favourite":
		//fixme: read from jwt
		total, err = h.Store.BeatmapSet().SetFavourite(103, uint(beatmapsetID))
	case "unfavourite":
		//fixme: read from jwt
		total, err = h.Store.BeatmapSet().SetUnFavourite(103, uint(beatmapsetID))
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	return c.JSON(http.StatusOK, struct {
		FavouriteCount uint `json:"favourite_count"`
	}{
		total,
	})
}
