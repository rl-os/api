package reqest_context

import (
	"context"
	myctx "github.com/deissh/osu-lazer/ayako/ctx"
	"github.com/labstack/echo/v4"
)

// GlobalMiddleware create new request reqest_context with all information about caller
func GlobalMiddleware(ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			if userId, ok := c.Get("current_user_id").(uint); ok {
				ctx = myctx.Pipe(ctx, myctx.SetUserID(userId))
			}

			if token, ok := c.Get("current_user_token").(string); ok {
				ctx = myctx.Pipe(ctx, myctx.SetUserToken(token))
			}

			if id := req.Header.Get(echo.HeaderXRequestID); id != "" {
				ctx = myctx.Pipe(ctx, myctx.SetRequestID(id))
			}

			c.Set("context", ctx)

			return next(c)
		}
	}
}
