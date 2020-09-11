package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/entity/request"
	"net/http"
	"strconv"
)

type BeatmapHandlers struct {
	App *app.App
}

func (h *BeatmapHandlers) Show(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	beatmapID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	beatmaps, err := h.App.Store.Beatmap().Get(ctx, uint(beatmapID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	return c.JSON(200, beatmaps)
}

func (h *BeatmapHandlers) Lookup(c echo.Context) (err error) {
	ctx, _ := c.Get("context").(context.Context)

	params := request.BeatmapLookup{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	var beatmap *entity.SingleBeatmap
	if params.CheckSum != "" {
		// todo: search by md5
		beatmap, err = h.App.Store.Beatmap().Get(ctx, params.Id)
	}

	if beatmap == nil && params.Id != 0 {
		beatmap, err = h.App.Store.Beatmap().Get(ctx, params.Id)
	}

	if beatmap == nil && params.Filename != "" {
		// todo: search by filename
		beatmap, err = h.App.Store.Beatmap().Get(ctx, params.Id)
	}

	if err != nil || beatmap == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	return c.JSON(200, beatmap)
}

func (h *BeatmapHandlers) Scores(c echo.Context) (err error) {
	params := request.GetBeatmapScores{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
