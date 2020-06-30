package permission

import (
	"context"
	"errors"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// keyFromHeader returns a `keyExtractor` that extracts key from the request header.
func keyFromHeader(header string) func(echo.Context) (string, error) {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		c.Logger().Info(auth)
		if auth == "" {
			return "", errors.New("missing key in request header")
		}
		if header == echo.HeaderAuthorization {
			l := len("Bearer")
			if len(auth) > l+1 && auth[:l] == "Bearer" {
				return auth[l+1:], nil
			}
			return "", errors.New("invalid key in the request header")
		}
		return auth, nil
	}
}

// GlobalMiddleware check access_token
func GlobalMiddleware(store store.Store, ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			extractor := keyFromHeader(echo.HeaderAuthorization)

			// check token and write to reqest_context if user send one
			if key, err := extractor(c); err == nil {
				token, err := store.OAuth().ValidateToken(ctx, key)
				if err != nil {
					// todo
					return next(c)
				}

				if err = store.User().UpdateLastVisit(ctx, token.UserID); err != nil {
					log.Error().Err(err).Msg("updating last visit")
				}

				c.Set("current_user_id", token.UserID)
				c.Set("current_user_token", token.AccessToken)
			}

			return next(c)
		}
	}
}
