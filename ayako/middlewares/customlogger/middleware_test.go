package customlogger

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	e := echo.New()

	got := Middleware()
	assert.NotNil(t, got)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderXRequestID, "someRequestId")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	middleware := got(func(context echo.Context) error {
		return echo.ErrServiceUnavailable
	})

	// it must just log and return original errors if exist
	err := middleware(c)
	assert.Equal(t, err, echo.ErrServiceUnavailable)
}
