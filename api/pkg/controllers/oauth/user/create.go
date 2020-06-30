package user

import (
	"github.com/deissh/rl/api/pkg"
	userService "github.com/deissh/rl/api/pkg/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// CreateTokenRequestData contain incoming data with user credentials
type CreateUserRequestData struct {
	Username string `json:"username" form:"user[username]" validate:"required"`
	Email    string `json:"email" form:"user[user_email]" validate:"required,email"`
	Password string `json:"password" form:"user[password]" validate:"required"`
}

// CreateUserHandler create new access_token and refresh_token pare
func CreateUserHandler(c echo.Context) (err error) {
	params := new(CreateUserRequestData)
	if err := c.Bind(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_params_error", "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	user, err := userService.Register(params.Username, params.Email, params.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
