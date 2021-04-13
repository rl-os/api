package api

import (
	"context"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
	"github.com/rl-os/api/entity/request"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/pkg/validator"
	"github.com/rs/zerolog"
)

type CurrentUserController struct {
	App       *app.App
	Logger    *zerolog.Logger
	Validator *validator.Inst
}

var providerMeSet = wire.NewSet(
	NewCurrentUserController,
)

func NewCurrentUserController(
	app *app.App,
	logger *zerolog.Logger,
	validator *validator.Inst,
) *CurrentUserController {
	return &CurrentUserController{
		app,
		logger,
		validator,
	}
}

// Create new user
//
// @Router /users [post]
// @Summary Create new user
// @Description get string by ID
// @Param payload body request.CreateUser true "JSON payload"
//
// @Success 200 {object} entity.User
// @Success 400 {object} errors.ResponseFormat
func (h *CurrentUserController) Create(c echo.Context) error {
	params := &request.CreateUser{}

	if err := c.Bind(params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	if err := h.Validator.Struct(params); err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	ctx, _ := c.Get("context").(context.Context)

	user, err := h.App.Store.User().Create(ctx, params.Username, params.Email, params.Password)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

// Me (current user)
//
// @Router /api/v2/me/{mode} [get]
// @Tags Current user
// @Summary Return current user
// @Security OAuth2
// @Param mode path string false "game mod"
//
// @Success 200 {object} entity.User
// @Success 400 {object} errors.ResponseFormat
func (h *CurrentUserController) Me(c echo.Context) error {
	ctx, _ := c.Get("context").(context.Context)
	mode := c.Param("mode")

	userId, err := myctx.GetUserID(ctx)
	if err != nil {
		return errors.New("request_params", 400, "Invalid params")
	}

	user, err := h.App.GetUser(ctx, userId, mode)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
