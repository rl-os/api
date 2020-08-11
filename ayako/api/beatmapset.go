package api

import (
	"context"
	"github.com/deissh/rl/ayako/app"
	myctx "github.com/deissh/rl/ayako/ctx"
	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BeatmapSetHandlers struct {
	*app.App
}

func (h *BeatmapSetHandlers) Get(c echo.Context) error {
	beatmapsetID, err := strconv.ParseUint(c.Param("beatmapset"), 10, 32)
	if err != nil {
		return errors.New("request_params", 400, "Invalid beatmapset id")
	}

	ctx := context.Background()

	userId, ok := c.Get("current_user_id").(uint)
	if ok {
		ctx = context.WithValue(context.Background(), "current_user_id", userId)
	}

	beatmaps, err := h.Store().BeatmapSet().Get(ctx, uint(beatmapsetID))
	if err != nil {
		return err
	}

	return c.JSON(200, beatmaps)
}

func (h *BeatmapSetHandlers) Lookup(c echo.Context) (err error) {
	params := struct {
		Id uint `query:"beatmap_id"`
	}{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	beatmap, err := h.Store().Beatmap().Get(ctx, params.Id)
	if err != nil {
		return err
	}

	beatmapSet, err := h.Store().BeatmapSet().Get(ctx, uint(beatmap.Beatmapset.ID))
	if err != nil {
		return err
	}

	return c.JSON(200, beatmapSet)
}

func (h *BeatmapSetHandlers) Search(c echo.Context) (err error) {
	// todo: this
	params := struct {
		Query    string `json:"q" query:"q"`
		Mode     int    `json:"m" query:"m"`
		Status   string `json:"s" query:"s"`
		Genre    string `json:"g" query:"g"`
		Language string `json:"l" query:"l"`
		Sort     string `json:"sort" query:"sort"`
	}{}
	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	//ctx := context.Background()
	//
	//userId, ok := c.Get("current_user_id").(uint)
	//if ok {
	//	ctx = context.WithValue(context.Background(), "current_user_id", userId)
	//}

	//beatmapSets, err := h.GetStore.BeatmapSet().Get(ctx, 1118896)
	//if err != nil {
	//	return err
	//}

	return c.JSON(200, struct {
		Beatmapsets           *[]entity.BeatmapSearch `json:"beatmapsets"`
		RecommendedDifficulty float32                 `json:"recommended_difficulty"`
		Error                 error                   `json:"error"`
		Total                 uint                    `json:"total"`
	}{
		&[]entity.BeatmapSearch{},
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
		return errors.New("request_params", 400, "Invalid beatmapset id")
	}

	if err := c.Bind(&params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return err
	}

	var total uint
	switch params.Action {
	case "favourite":
		total, err = h.Store().BeatmapSet().SetFavourite(ctx, userId, uint(beatmapsetID))
	case "unfavourite":
		total, err = h.Store().BeatmapSet().SetUnFavourite(ctx, userId, uint(beatmapsetID))
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		FavouriteCount uint `json:"favourite_count"`
	}{
		total,
	})
}
