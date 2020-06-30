package oauth

import (
	"github.com/deissh/rl/api/pkg"
	"github.com/deissh/rl/api/pkg/controllers/oauth/token"
	"github.com/deissh/rl/api/pkg/controllers/oauth/user"
	oauthService "github.com/deissh/rl/api/pkg/services/oauth"
	"github.com/labstack/echo/v4"
)

func empty(_ echo.Context) (err error) {
	return pkg.NewHTTPError(400, "invalid_credentials", "The user credentials were incorrect.")
}

func scopes(c echo.Context) (err error) {
	return c.JSON(200, oauthService.Scopes)
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *echo.Group) {
	// OAuth method
	// like create and revoke token
	oauth := r.Group("/oauth")
	{
		// Requesting Tokens
		oauth.POST("/token", token.CreateTokenHandler) // https://laravel.com/docs/5.7/passport#refreshing-tokens

		// Managing Tokens
		oauth.GET("/scopes", scopes)

		// Managing Clients
		oauth.GET("/clients", empty)
		oauth.POST("/clients", empty)
		oauth.PUT("/clients/:client", empty)
	}

	// User registration and password recover
	r.POST("/users", user.CreateUserHandler)
}
