package token

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

// CreateTokenRequestData contain incoming data with user credentials
type CreateTokenRequestData struct {
	GrantType    string `json:"grant_type" form:"grant_type" validate:"required"`
	ClientID     string `json:"client_id" form:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" form:"client_secret" validate:"required"`
	Scope        string `json:"scope" form:"scope" validate:"required"`

	// grant_type == password
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`

	// grant_type == refresh_token
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

// CreateTokenResponseData response struct that contain new access_token and refresh_token
type CreateTokenResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// CreateTokenHandler create new access_token and refresh_token pare
func CreateTokenHandler(c echo.Context) (err error) {
	params := new(CreateTokenRequestData)

	if err := c.Bind(params); err != nil {
		return c.JSON(400, pkg.ErrorResponse{
			Error:            "params_error",
			ErrorDescription: "Failed binding params",
			Message:          err.Error(),
		})
	}

	if err := validator.New().Struct(params); err != nil {
		return c.JSON(400, pkg.ErrorResponse{
			Error:            "validate_error",
			ErrorDescription: "Failed validate",
			Message:          err.Error(),
		})
	}

	token := services.OAuthToken{
		UserId:       1,
		AccessToken:  "",
		RefreshToken: "",
	}

	_, err = pkg.Db.NamedExec(
		"INSERT INTO oauth_token (user_id, access_token, refresh_token, scopes) VALUES (:user_id, :access_token, :refresh_token)",
		token)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Creating token")
	}

	return c.JSON(http.StatusOK, CreateTokenResponseData{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    "bearer",
		ExpiresIn:    0,
	})
}
