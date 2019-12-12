package user

import (
	userService "github.com/deissh/osu-api-server/pkg/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateUserRequestData contain incoming data with user credentials
type CreateUserRequestData struct {
	Username string `json:"username" form:"user[username]" validate:"required"`
	Email string `json:"email" form:"user[user_email]" validate:"required,email"`
	Password string `json:"password" form:"user[password]" validate:"required"`
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

