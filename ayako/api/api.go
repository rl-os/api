package api

import (
	"github.com/deissh/rl/ayako/middlewares/permission"
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
)

func New(store store.Store, router *echo.Echo) {
	signup := router.Group("/users")
	{
		h := RegistrationHandlers{store}
		signup.POST("", h.Create)
	}

	oauth := router.Group("/oauth")
	{
		h := OAuthTokenHandlers{store}
		oauth.POST("/token", h.Create)
	}

	v2 := router.Group("/api/v2")
	{
		// Health status
		// хз для чего я пишу в каждой версии свой пинг
		v2.GET("/ping", echo.MethodNotAllowedHandler)

		// === Me ===
		me := v2.Group("/me", permission.MustLogin)
		{
			h := MeHandlers{store}
			me.GET("/", h.Me)
			me.GET("/:mode", h.Me)
			me.GET("/me/download-quota-check", echo.MethodNotAllowedHandler)
		}

		// === Friends ===
		friends := v2.Group("/friends", permission.MustLogin)
		{
			h := FriendHandlers{store}
			friends.GET("", h.GetAll)
			friends.PUT("", h.Add)
			friends.DELETE("", h.Remove)
		}

		// === Users ===
		users := v2.Group("/users")
		{
			users.GET("/:user/kudosu", echo.MethodNotAllowedHandler)
			users.GET("/:user/scores/:type", echo.MethodNotAllowedHandler)
			users.GET("/:user/beatmapsets/:type", echo.MethodNotAllowedHandler)
			users.GET("/:user/recent_activity", echo.MethodNotAllowedHandler)
			users.GET("/:user/:mode", echo.MethodNotAllowedHandler)
			users.GET("/:user/", echo.MethodNotAllowedHandler)
		}

		// === Beatmaps ===
		bmaps := v2.Group("/beatmaps")
		{
			h := BeatmapHandlers{store}
			bmaps.GET("/beatmaps/lookup", h.Lookup)
			bmaps.GET("/beatmaps/:beatmap", h.Show)
			bmaps.GET("/beatmaps/:beatmap/scores", h.Scores)
		}

		// === Beatmapsets ===
		bmsets := v2.Group("/beatmapsets")
		{
			h := BeatmapSetHandlers{store}
			bmsets.GET("/lookup", h.Lookup)
			bmsets.GET("/search/:filters", h.Search)
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
			chat.POST("/new", echo.MethodNotAllowedHandler)
			chat.GET("/updates", echo.MethodNotAllowedHandler)
			chat.GET("/presence", echo.MethodNotAllowedHandler) // ???
			chat.GET("/channels", echo.MethodNotAllowedHandler)
			chat.GET("/channels/joined", echo.MethodNotAllowedHandler)
			chat.GET("/channels/:channel/messages", echo.MethodNotAllowedHandler)
			chat.POST("/channels/:channel/messages", echo.MethodNotAllowedHandler)
			chat.PUT("/channels/:channel/users/:user", echo.MethodNotAllowedHandler)
			chat.DELETE("/channels/:channel/users/:user", echo.MethodNotAllowedHandler)
			chat.PUT("/channels/:channel/mark-as-read/:message", echo.MethodNotAllowedHandler)
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

		// === Misc ===
		v2.POST("/reports", echo.MethodNotAllowedHandler)
		v2.GET("/changelog", echo.MethodNotAllowedHandler)
		v2.GET("/changelog/:changelog", echo.MethodNotAllowedHandler)
	}
}
