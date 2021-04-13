package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/pkg/validator"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

var providerBeatmapSet = wire.NewSet(
	NewBeatmapController,
)

type BeatmapController struct {
	App       *app.App
	Logger    *zerolog.Logger
	Validator *validator.Inst
}

func NewBeatmapController(
	app *app.App,
	logger *zerolog.Logger,
	validator *validator.Inst,
) *BeatmapController {
	return &BeatmapController{
		app,
		logger,
		validator,
	}
}

// Get beatmap by id
//
// @Router /api/v2/beatmaps/{id} [get]
// @Tags Beatmap
// @Summary Return beatmap by id
// @Param id path string true "beatmap id"
//
// @Success 200 {object} entity.SingleBeatmap
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapController) Get(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	beatmapID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	beatmaps, err := h.App.GetBeatmap(ctx, uint(beatmapID))
	if err != nil {
		return err
	}

	return c.JSON(200, beatmaps)
}

// Lookup beatmap by id, checksum, filename
//
// @Router /api/v2/beatmaps/lookup [get]
// @Tags Beatmap
// @Summary Lookup beatmap by id, checksum, filename
// @Param id query integer false "beatmap id"
// @Param checksum query string false "beatmap file md5"
// @Param filename query string false "beatmap filename (legacy)"
//
// @Success 200 {object} entity.SingleBeatmap
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapController) Lookup(c echo.Context) (err error) {
	ctx, _ := c.Get("context").(context.Context)

	params := request.BeatmapLookup{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	beatmap, err := h.App.LookupBeatmap(ctx, params.Id, params.CheckSum, params.Filename)
	if err != nil {
		return err
	}

	return c.JSON(200, beatmap)
}

// Scores submitted to selected beatmap
//
// @Router /api/v2/beatmaps/{id}/scores [get]
// @Tags Beatmap
// @Summary Scores submitted to selected beatmap
// @Param beatmap path string true "beatmap id"
// @Param type query string false "score type"
// @Param mode query string false "osu! game mode (std, mania, ctb and etc)"
//
// @Success 200 {object} entity.SingleBeatmap
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapController) Scores(c echo.Context) (err error) {
	params := request.GetBeatmapScores{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	_, err = strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid beatmap id")
	}

	return c.JSON(http.StatusOK, nil)
}
