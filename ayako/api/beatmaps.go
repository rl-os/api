package api

import (
	"context"
	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/middlewares/permission"
	"github.com/deissh/rl/ayako/store"
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
	api.Beatmaps.GET("/:id/scores", echo.NotFoundHandler, permission.MustLogin)
	api.Beatmaps.GET("/lookup", handlers.Lookup)
}

func (h *BeatmapHandlers) Show(c echo.Context) error {
	beatmapID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	beatmaps, err := h.Store.Beatmap().Get(ctx, uint(beatmapID))
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

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	var beatmap *entity.SingleBeatmap
	if params.CheckSum != "" {
		// todo: search by md5
		beatmap, err = h.Store.Beatmap().Get(ctx, params.Id)
	}

	if beatmap == nil && params.Id != 0 {
		beatmap, err = h.Store.Beatmap().Get(ctx, params.Id)
	}

	if beatmap == nil && params.Filename != "" {
		// todo: search by filename
		beatmap, err = h.Store.Beatmap().Get(ctx, params.Id)
	}

	if err != nil || beatmap == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	return c.JSON(200, beatmap)
}

func (h *BeatmapHandlers) Scores(c echo.Context) (err error) {
	params := struct {
		Type string `query:"type"`
		Mode string `query:"mode"`
	}{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
