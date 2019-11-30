package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func empty(c echo.Context) (err error) {
	return c.JSON(200, gin.H{
		"message": "I'm alive!",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *echo.Group) {
	v2 := r.Group("/v2")
	{
		// Health status
		// хз для чего я пишу в каждой версии свой пинг
		v2.GET("/ping", empty)

		// === Me ===
		v2.GET("/me/:mode", empty)
		v2.GET("/me/download-quota-check", empty)

		// === Friends ===
		v2.GET("/friends", empty)

		// === Users ===
		v2.GET("/users/:user/kudosu", empty)
		v2.GET("/users/:user/scores/:type", empty)
		v2.GET("/users/:user/beatmapsets/:type", empty)
		v2.GET("/users/:user/recent_activity", empty)
		v2.GET("/users/:user/:mode", empty)

		// === Beatmaps ===
		v2.GET("/beatmaps/lookup", empty)
		v2.GET("/beatmaps/:beatmap", empty)
		v2.GET("/beatmaps/:beatmap/scores", empty)
		v2.GET("/beatmapsets/lookup", empty)
		v2.GET("/beatmapsets/search/:filters", empty)
		v2.GET("/beatmapsets/:beatmapset", empty)
		v2.GET("/beatmapsets/:beatmapset/download", empty)
		v2.GET("/beatmapsets/:beatmapset/favourites", empty)

		// === Scores ===
		v2.GET("/scores/:mode/:score/download", empty)

		// === Rooms ===
		v2.POST("/rooms", empty)
		v2.GET("/rooms/:room", empty)
		v2.PUT("/rooms/:room/users/:user", empty)
		v2.DELETE("/rooms/:room/users/:user", empty)
		v2.GET("/rooms/:room/leaderboard", empty)
		v2.POST("/rooms/:room/playlist/:playlist/scores", empty)
		v2.PUT("/rooms/:room/playlist/:playlist/scores/:score", empty)

		// === Chats ===
		v2.POST("/chat/new", empty)
		v2.GET("/chat/updates", empty)
		v2.GET("/chat/presence", empty) // ???
		v2.GET("/chat/channels", empty)
		v2.GET("/chat/channels/:channel/messages", empty)
		v2.POST("/chat/channels/:channel/messages", empty)
		v2.PUT("/chat/channels/:channel/users/:user", empty)
		v2.DELETE("/chat/channels/:channel/users/:user", empty)
		v2.PUT("/chat/channels/:channel/mark-as-read/:message", empty)

		// === Comments ===
		v2.GET("/comments", empty)
		v2.POST("/comments", empty)
		v2.GET("/comments/:comment", empty)
		v2.PUT("/comments/:comment", empty)
		v2.PATCH("/comments/:comment", empty)
		v2.DELETE("/comments/:comment", empty)
		v2.POST("/comments/:comment/vote", empty)
		v2.DELETE("/comments/:comment/vote", empty)

		// === Notifications ===
		v2.GET("/notifications", empty)
		v2.POST("/notifications/mark-read", empty)

		// === Misc ===
		v2.POST("/reports", empty)
		v2.GET("/changelog", empty)
		v2.GET("/changelog/:changelog", empty)
	}
}
