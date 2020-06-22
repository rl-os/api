package api

import (
	"context"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/middlewares/permission"
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
	api.BeatmapSets.POST("/:beatmapset/favourites", handlers.Favourite, permission.MustLogin)
	api.BeatmapSets.GET("/:beatmapset/download", handlers.Get, permission.MustLogin)
	api.BeatmapSets.GET("/lookup", handlers.Lookup)
	api.BeatmapSets.GET("/search", handlers.Search)
}

func (h *BeatmapSetHandlers) Get(c echo.Context) error {
	beatmapsetID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmapset id")
	}

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	beatmaps, err := h.Store.BeatmapSet().Get(ctx, uint(beatmapsetID))
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

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	beatmap, err := h.Store.Beatmap().Get(ctx, params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	beatmapSet, err := h.Store.BeatmapSet().Get(ctx, uint(beatmap.Beatmapset.ID))
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

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	beatmapSets, err := h.Store.BeatmapSet().Get(ctx, 1118896)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "BeatmapSet not found")
	}

	return c.JSON(200, struct {
		Beatmapsets           *[]entity.BeatmapSetFull `json:"beatmapsets"`
		RecommendedDifficulty float32                  `json:"recommended_difficulty"`
		Error                 error                    `json:"errors"`
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

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	var total uint
	switch params.Action {
	case "favourite":
		total, err = h.Store.BeatmapSet().SetFavourite(ctx, userId, uint(beatmapsetID))
	case "unfavourite":
		total, err = h.Store.BeatmapSet().SetUnFavourite(ctx, userId, uint(beatmapsetID))
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal errors")
	}

	return c.JSON(http.StatusOK, struct {
		FavouriteCount uint `json:"favourite_count"`
	}{
		total,
	})
}
