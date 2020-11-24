//go:generate swag init -o ../docs -g ./api/api.go --dir .././

package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/middlewares/permission"
)

// New api endpoint
//
// @title osu!lazer API
// @version 2.0
// @description This is a simple server.
// @host localhost:2400
// @BasePath /
//
// @contact.name RL GitHub
// @contact.url https://github.com/rl-os/api
//
// @securitydefinitions.oauth2.password OAuth2
// @in request-body
// @tokenUrl /oauth/token
// @scope.* Grants all access
// @scope.user.* Grants access as user (without system and admin api)
// @scope.admin.* Grants access as admin
// @scope.sys.* Grants access as system user (for example chatbot, worker and etc)
func New(app *app.App, root *echo.Group) {
	signup := root.Group("/users")
	{
		h := RegistrationHandlers{app}
		signup.POST("", h.Create)
	}

	// osu! lazer oauth2 support
	// HACK: remove if possible
	oauth := root.Group("/oauth")
	{
		h := OAuthTokenHandlers{app}
		oauth.POST("/token", h.Create)
	}

	v2 := root.Group("/api/v2")
	{
		v2.GET("/ping", echo.MethodNotAllowedHandler)

		// === Me ===
		me := v2.Group("/me", permission.MustLogin)
		{
			h := MeHandlers{app}
			me.GET("/", h.Me)
			me.GET("/:mode", h.Me)
			me.GET("/me/download-quota-check", echo.MethodNotAllowedHandler)
		}

		// === Friends ===
		friends := v2.Group("/friends", permission.MustLogin)
		{
			h := FriendHandlers{app}
			friends.GET("", h.GetAll)
			friends.PUT("", h.Add)
			friends.DELETE("", h.Remove)
		}

		// === Users ===
		users := v2.Group("/users")
		{
			h := UsersHandlers{app}
			users.GET("/:user/kudosu", echo.MethodNotAllowedHandler)
			users.GET("/:user/scores/:type", echo.MethodNotAllowedHandler)
			users.GET("/:user/beatmapsets/:type", echo.MethodNotAllowedHandler)
			users.GET("/:user/recent_activity", echo.MethodNotAllowedHandler)
			users.GET("/:user/:mode", h.Get)
			users.GET("/:user/", h.Get)
			users.GET("/:user", h.Get)
		}

		// === Beatmaps ===
		bmaps := v2.Group("/beatmaps")
		{
			h := BeatmapHandlers{app}
			bmaps.GET("/lookup", h.Lookup)
			bmaps.GET("/:beatmap", h.Get)
			bmaps.GET("/:beatmap/scores", h.Scores)
		}

		// === Beatmapsets ===
		bmsets := v2.Group("/beatmapsets")
		{
			h := BeatmapSetHandlers{app}
			bmsets.GET("/lookup", h.Lookup)
			bmsets.GET("/search", h.Search)
			bmsets.GET("/search/:filters", h.Search) // ???
			bmsets.GET("/:beatmapset", h.Get)
			bmsets.GET("/:beatmapset/download", echo.MethodNotAllowedHandler, permission.MustLogin)
			bmsets.POST("/:beatmapset/favourites", h.Favourite, permission.MustLogin)
		}

		// === Scores ===
		scores := v2.Group("/scores")
		{
			scores.GET("/:mode/:score/download", echo.MethodNotAllowedHandler)
		}

		// === Rooms ===
		rooms := v2.Group("/rooms")
		{
			rooms.POST("", echo.MethodNotAllowedHandler)
			rooms.GET("/:room", echo.MethodNotAllowedHandler)
			rooms.PUT("/:room/users/:user", echo.MethodNotAllowedHandler)
			rooms.DELETE("/:room/users/:user", echo.MethodNotAllowedHandler)
			rooms.GET("/:room/leaderboard", echo.MethodNotAllowedHandler)
			rooms.POST("/:room/playlist/:playlist/scores", echo.MethodNotAllowedHandler)
			rooms.PUT("/:room/playlist/:playlist/scores/:score", echo.MethodNotAllowedHandler)
		}

		// === Chats ===
		chat := v2.Group("/chat", permission.MustLogin)
		{
			h := ChatHandlers{app}
			chat.POST("/new", h.NewPm)
			chat.GET("/updates", h.Updates)
			chat.GET("/presence", echo.MethodNotAllowedHandler) // ???
			chat.GET("/channels", h.GetAll)
			chat.GET("/channels/joined", h.GetJoined)
			chat.GET("/channels/:channel/messages", h.Messages)
			chat.POST("/channels/:channel/messages", h.Send)
			chat.PUT("/channels/:channel/users/:user", h.Join)
			chat.DELETE("/channels/:channel/users/:user", h.Leave)
			chat.PUT("/channels/:channel/mark-as-read/:message", echo.MethodNotAllowedHandler) // todo
		}

		// === Comments ===
		comments := v2.Group("/comments")
		{
			comments.GET("/", echo.MethodNotAllowedHandler)
			comments.POST("/", echo.MethodNotAllowedHandler)
			comments.GET("/:comment", echo.MethodNotAllowedHandler)
			comments.PUT("/:comment", echo.MethodNotAllowedHandler)
			comments.PATCH("/:comment", echo.MethodNotAllowedHandler)
			comments.DELETE("/:comment", echo.MethodNotAllowedHandler)
			comments.POST("/:comment/vote", echo.MethodNotAllowedHandler)
			comments.DELETE("/:comment/vote", echo.MethodNotAllowedHandler)
		}

		// === Notifications ===
		notif := v2.Group("/Notifications", permission.MustLogin)
		{
			notif.GET("/", echo.MethodNotAllowedHandler)
			notif.POST("/mark-read", echo.MethodNotAllowedHandler)
		}

		// === Private OAuth configurations
		oauth := v2.Group("/oauth")
		{
			client := oauth.Group("/client")
			{
				h := OAuthClientHandlers{app}
				client.POST("", h.Create, permission.MustLogin)
			}

			token := oauth.Group("/token")
			{
				h := OAuthTokenHandlers{app}
				token.POST("", h.Create)
			}
		}

		// === Misc ===
		v2.POST("/reports", echo.MethodNotAllowedHandler)
		v2.GET("/changelog", echo.MethodNotAllowedHandler)
		v2.GET("/changelog/:changelog", echo.MethodNotAllowedHandler)
	}
}
