package permission

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
func GlobalMiddleware(jwtSecret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			extractor := keyFromHeader(echo.HeaderAuthorization)

			// check token and write to reqest_context if user send one
			if key, err := extractor(c); err == nil {
				claims := jwt.MapClaims{}
				token, err := jwt.ParseWithClaims(key, &claims, func(token *jwt.Token) (interface{}, error) {
					return jwtSecret, nil
				})
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
				}

				userId, err := strconv.ParseUint(claims["sub"].(string), 10, 32)
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
				}

				c.Set("current_user_id", uint(userId))
				c.Set("current_user_token", token)
			}

			return next(c)
		}
	}
}
