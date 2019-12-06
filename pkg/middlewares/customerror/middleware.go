package customerror

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

// ErrorResponse default struct for errors
type ErrorResponse struct {
	// contain short, unique error
	Error string `json:"error"`
	// full description about error
	ErrorDescription string `json:"error_description"`
	// addition information
	Message string `json:"message"`
}

// CustomHTTPErrorHandler transform GoLang error to JSON response
func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Size > 0 { return }
	
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// todo: getting locale and translate error message

	err = c.JSON(code, ErrorResponse{
		Error:            err.Error(), // todo: this
		ErrorDescription: err.Error(), // todo: this
		Message:          http.StatusText(code), // todo: this
	})
	if err != nil { log.Error().Err(err) }
}
