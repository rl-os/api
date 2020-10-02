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

// Show beatmap by id
//
// @Router /api/v2/beatmaps/{beatmap_id} [get]
// @Tags Beatmap
// @Summary Return beatmap by id
// @Param beatmap_id path string true "beatmap id"
//
// @Success 200 {object} entity.SingleBeatmap
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapHandlers) Show(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	beatmapID, err := strconv.ParseUint(c.Param("beatmap"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	beatmaps, err := h.App.Store.Beatmap().Get(ctx, uint(beatmapID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Beatmap not found")
	}

	return c.JSON(200, beatmaps)
}

// Lookup beatmap by id, checksum, filename
//
// @Router /api/v2/beatmaps/lookup [get]
// @Tags Beatmap
// @Summary Lookup beatmap by id, checksum, filename
// @Param id query string false "beatmap id"
// @Param checksum query string false "beatmap file md5"
// @Param filename query string false "beatmap filename (legacy)"
//
// @Success 200 {object} entity.SingleBeatmap
// @Success 400 {object} errors.ResponseFormat
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

// Scores submitted to selected beatmap
//
// @Router /api/v2/beatmaps/{beatmap}/scores [get]
// @Tags Beatmap
// @Summary Scores submitted to selected beatmap
// @Param beatmap_id path string true "beatmap id"
// @Param type query string false "score type"
// @Param mode query string false "osu! game mode (std, mania, ctb and etc)"
//
// @Success 200 {object} entity.SingleBeatmap
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapHandlers) Scores(c echo.Context) (err error) {
	params := request.GetBeatmapScores{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	_, err = strconv.ParseUint(c.Param("beatmap"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	return c.JSON(http.StatusOK, nil)
}
