package user

import (
	userService "github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateTokenRequestData contain incoming data with user credentials
type CreateUserRequestData struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// CreateTokenHandler create new access_token and refresh_token pare
func CreateTokenHandler(c echo.Context) (err error) {
	params := new(CreateUserRequestData)
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed validate")
	}

	user, err := userService.Register(params.Username, params.Email, params.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

