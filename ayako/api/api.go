package api

import (
	"github.com/deissh/rl/ayako/middlewares/permission"
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Beatmaps    *echo.Group
	BeatmapSets *echo.Group
	Me          *echo.Group
	Friend      *echo.Group
}

func New(store store.Store, prefix *echo.Group) {
	v2 := prefix.Group("/v2")
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
			bmsets.GET("/:beatmapset/favourites", h.Favourite, permission.MustLogin)
		}

		// === Scores ===
		v2.GET("/scores/:mode/:score/download", echo.MethodNotAllowedHandler)

		// === Rooms ===
		v2.POST("/rooms", echo.MethodNotAllowedHandler)
		v2.GET("/rooms/:room", echo.MethodNotAllowedHandler)
		v2.PUT("/rooms/:room/users/:user", echo.MethodNotAllowedHandler)
		v2.DELETE("/rooms/:room/users/:user", echo.MethodNotAllowedHandler)
		v2.GET("/rooms/:room/leaderboard", echo.MethodNotAllowedHandler)
		v2.POST("/rooms/:room/playlist/:playlist/scores", echo.MethodNotAllowedHandler)
		v2.PUT("/rooms/:room/playlist/:playlist/scores/:score", echo.MethodNotAllowedHandler)

		// === Chats ===
		v2Chat := v2.Group("/chat", permission.MustLogin)
		{
			v2Chat.POST("/new", echo.MethodNotAllowedHandler)
			v2Chat.GET("/updates", echo.MethodNotAllowedHandler)
			v2Chat.GET("/presence", echo.MethodNotAllowedHandler) // ???
			v2Chat.GET("/channels", echo.MethodNotAllowedHandler)
			v2Chat.GET("/channels/joined", echo.MethodNotAllowedHandler)
			v2Chat.GET("/channels/:channel/messages", echo.MethodNotAllowedHandler)
			v2Chat.POST("/channels/:channel/messages", echo.MethodNotAllowedHandler)
			v2Chat.PUT("/channels/:channel/users/:user", echo.MethodNotAllowedHandler)
			v2Chat.DELETE("/channels/:channel/users/:user", echo.MethodNotAllowedHandler)
			v2Chat.PUT("/channels/:channel/mark-as-read/:message", echo.MethodNotAllowedHandler)
		}

		// === Comments ===
		v2Comments := v2.Group("/comments")
		{
			v2Comments.GET("/", echo.MethodNotAllowedHandler)
			v2Comments.POST("/", echo.MethodNotAllowedHandler)
			v2Comments.GET("/:comment", echo.MethodNotAllowedHandler)
			v2Comments.PUT("/:comment", echo.MethodNotAllowedHandler)
			v2Comments.PATCH("/:comment", echo.MethodNotAllowedHandler)
			v2Comments.DELETE("/:comment", echo.MethodNotAllowedHandler)
			v2Comments.POST("/:comment/vote", echo.MethodNotAllowedHandler)
			v2Comments.DELETE("/:comment/vote", echo.MethodNotAllowedHandler)
		}

		// === Notifications ===
		v2Notif := v2.Group("/Notifications", permission.MustLogin)
		{
			v2Notif.GET("/", echo.MethodNotAllowedHandler)
			v2Notif.POST("/mark-read", echo.MethodNotAllowedHandler)
		}

		// === Misc ===
		v2.POST("/reports", echo.MethodNotAllowedHandler)
		v2.GET("/changelog", echo.MethodNotAllowedHandler)
		v2.GET("/changelog/:changelog", echo.MethodNotAllowedHandler)
	}
}
