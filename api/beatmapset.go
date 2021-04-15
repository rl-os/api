package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/services/validator"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

var providerBeatmapsetSet = wire.NewSet(
	NewBeatmapSetController,
)

type BeatmapSetController struct {
	UseCase *app.BeatmapSetUseCase

	Logger    *zerolog.Logger
	Validator *validator.Inst
}

func NewBeatmapSetController(
	useCase *app.BeatmapSetUseCase,
	logger *zerolog.Logger,
	validator *validator.Inst,
) *BeatmapSetController {
	return &BeatmapSetController{
		useCase,
		logger,
		validator,
	}
}

// Get beatmap by id
//
// @Router /api/v2/beatmapsets/{id} [get]
// @Tags Beatmapset
// @Summary Get beatmap by id
// @Param id path int true "beatmapset id"
//
// @Success 200 {object} entity.BeatmapSetFull
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapSetController) Get(c echo.Context) error {
	beatmapsetID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.New("request_params", 400, "Invalid beatmapset id")
	}

	ctx, _ := c.Get("context").(context.Context)

	beatmapset, err := h.UseCase.GetBeatmapset(ctx, uint(beatmapsetID))
	if err != nil {
		return err
	}

	return c.JSON(200, beatmapset)
}

// Lookup beatmapset by beatmap id
//
// @Router /api/v2/beatmapsets/lookup [get]
// @Tags Beatmapset
// @Summary Lookup beatmapset by beatmap id
// @Param beatmap_id query integer false "beatmap id"
//
// @Success 200 {object} entity.BeatmapSetFull
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapSetController) Lookup(c echo.Context) (err error) {
	params := request.BeatmapsetLookup{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	beatmapset, err := h.UseCase.LookupBeatmapset(ctx, params.Id)
	if err != nil {
		return err
	}

	return c.JSON(200, beatmapset)
}

// Search beatmapset
//
// @Router /api/v2/beatmapsets/search [get]
// @Tags Beatmapset
// @Summary Search Beatmapset
// @Param q query string false "query"
// @Param m query integer false "mode"
// @Param s query string false "status"
// @Param g query string false "genre"
// @Param l query string false "lang"
// @Param sort query string false "sort"
//
// @Success 200 {object} entity.BeatmapsetSearchResult
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapSetController) Search(c echo.Context) (err error) {
	params := request.BeatmapsetSearch{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	result, err := h.UseCase.SearchBeatmapset(ctx)
	if err != nil {
		return err
	}

	return c.JSON(200, result)
}

// Favourite beatmapset
//
// @Router /api/v2/beatmapsets/{id}/favourites [post]
// @Tags Beatmapset
// @Summary Favourite beatmapset
// @Security OAuth2
// @Param id path int true "beatmapset id"
// @Param action query string false "action" Enums(favourite, unfavourite)
//
// @Success 200 {object} object
// @Success 400 {object} errors.ResponseFormat
func (h *BeatmapSetController) Favourite(c echo.Context) (err error) {
	beatmapsetID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return errors.New("request_params", 400, "Invalid beatmapset id")
	}

	params := struct {
		Action string `query:"action" json:"action"`
	}{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	total, err := h.UseCase.FavouriteBeatmapset(ctx, params.Action, uint(beatmapsetID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		FavouriteCount uint `json:"favourite_count"`
	}{
		total,
	})
}
