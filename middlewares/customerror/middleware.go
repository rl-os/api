package customerror

import (
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/errors"
	"github.com/rs/zerolog/log"

	"net/http"
)

// CustomHTTPErrorHandler transform GoLang errors to JSON response
func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Size > 0 {
		return
	}

	code := http.StatusInternalServerError
	res := errors.ResponseFormat{
		ErrorID:          "internal",
		ErrorDescription: err.Error(),
		Message:          http.StatusText(code),
	}

	if he, ok := err.(*echo.HTTPError); ok {
		res.ErrorID = "internal"
		res.ErrorDescription = he.Error()
		res.Message = http.StatusText(he.Code)
		code = he.Code
	}

	// if custom errors
	if he, ok := err.(*errors.Error); ok {
		res = he.ResponseFormat()
		code = he.Code
	}

	// todo: getting locale and translate errors message

	err = c.JSON(code, res)
	if err != nil {
		log.Error().Err(err)
	}
}
