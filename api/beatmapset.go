package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"net/http"
	"strconv"
)

type BeatmapSetHandlers struct {
	App *app.App
}

func (h *BeatmapSetHandlers) Get(c echo.Context) error {
	beatmapsetID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
	if err != nil {
		return errors.New("request_params", 400, "Invalid beatmapset id")
	}

	ctx, _ := c.Get("context").(context.Context)

	beatmapset, err := h.App.GetBeatmapset(ctx, uint(beatmapsetID))
	if err != nil {
		return err
	}

	return c.JSON(200, beatmapset)
}

func (h *BeatmapSetHandlers) Lookup(c echo.Context) (err error) {
	params := request.BeatmapsetLookup{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	beatmapset, err := h.App.LookupBeatmapset(ctx, params.Id)
	if err != nil {
		return err
	}

	return c.JSON(200, beatmapset)
}

func (h *BeatmapSetHandlers) Search(c echo.Context) (err error) {
	params := request.BeatmapsetSearch{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	result, err := h.App.SearchBeatmapset(ctx)
	if err != nil {
		return err
	}

	return c.JSON(200, result)
}

func (h *BeatmapSetHandlers) Favourite(c echo.Context) (err error) {
	beatmapsetID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
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

	total, err := h.App.FavouriteBeatmapset(ctx, params.Action, uint(beatmapsetID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		FavouriteCount uint `json:"favourite_count"`
	}{
		total,
	})
}
