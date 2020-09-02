package reqest_context

import (
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	myctx "github.com/rl-os/api/ctx"
)

// GlobalMiddleware create new request reqest_context with all information about caller
func GlobalMiddleware(app *app.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			// context is growing very quickly
			// due to the fact that it is used directly
			// so create a copy
			copyCtx := app.Context

			if userId, ok := c.Get("current_user_id").(uint); ok {
				copyCtx = myctx.Pipe(copyCtx, myctx.SetUserID(userId))
			}

			if token, ok := c.Get("current_user_token").(string); ok {
				copyCtx = myctx.Pipe(copyCtx, myctx.SetUserToken(token))
			}

			if id := req.Header.Get(echo.HeaderXRequestID); id != "" {
				copyCtx = myctx.Pipe(copyCtx, myctx.SetRequestID(id))
			}

			c.Set("context", copyCtx)

			return next(c)
		}
	}
}
