package customerror

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

// CustomHTTPErrorHandler transform GoLang error to JSON response
func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Size > 0 { return }

	code := http.StatusInternalServerError	
	res := pkg.ErrorResponse{
		Error:            err.Error(),
		ErrorDescription: err.Error(),
		Message:          http.StatusText(code),
	}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		res.Message = http.StatusText(code)
	}

	// if custom error
	if he, ok := err.(*pkg.HTTPError); ok {
		res = he.Body
	}

	// todo: getting locale and translate error message

	err = c.JSON(code, res)
	if err != nil { log.Error().Err(err) }
}
