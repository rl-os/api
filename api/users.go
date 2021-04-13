package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/pkg/validator"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type UserController struct {
	App       *app.App
	Logger    *zerolog.Logger
	Validator *validator.Inst
}

var providerUserSet = wire.NewSet(
	NewUserController,
)

func NewUserController(
	app *app.App,
	logger *zerolog.Logger,
	validator *validator.Inst,
) *UserController {
	return &UserController{
		app,
		logger,
		validator,
	}
}

func (h *UserController) Get(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)

	userId, err := strconv.ParseUint(c.Param("user"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mode := c.Param("mode")

	user, err := h.App.GetUser(ctx, uint(userId), mode)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
