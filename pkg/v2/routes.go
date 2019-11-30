package v2

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'm alive!",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v2 := r.Group("/v2")
	{
		// Health status
		// хз для чего я пишу в каждой версии свой пинг
		v2.GET("/ping", ping)

		// === Me ===
		v2.GET("/me/:mode")
		v2.GET("/me/download-quota-check")

		// === Friends ===
		v2.GET("/friends")
		
		// === Users ===
		v2.GET("/users/:user/kudosu")
		v2.GET("/users/:user/scores/:type")
		v2.GET("/users/:user/beatmapsets/:type")
		v2.GET("/users/:user/recent_activity")
		v2.GET("/users/:user/:mode")

		// === Beatmaps ===
		v2.GET("/beatmaps/lookup")
		v2.GET("/beatmaps/:beatmap")
		v2.GET("/beatmaps/:beatmap/scores")
		v2.GET("/beatmapsets/lookup")
		v2.GET("/beatmapsets/search/:filters")
		v2.GET("/beatmapsets/:beatmapset")
		v2.GET("/beatmapsets/:beatmapset/download")
		v2.GET("/beatmapsets/:beatmapset/favourites")

		// === Scores ===
		v2.GET("/scores/:mode/:score/download")

		// === Rooms ===
		v2.POST("/rooms")
		v2.GET("/rooms/:room")
		v2.PUT("/rooms/:room/users/:user")
		v2.DELETE("/rooms/:room/users/:user")
		v2.GET("/rooms/:room/leaderboard")
		v2.POST("/rooms/:room/playlist/:playlist/scores")
		v2.PUT("/rooms/:room/playlist/:playlist/scores/:score")

		// === Chats ===
		v2.POST("/chat/new")
		v2.GET("/chat/updates")
		v2.GET("/chat/presence") // ???
		v2.GET("/chat/channels")
		v2.GET("/chat/channels/:channel/messages")
		v2.POST("/chat/channels/:channel/messages")
		v2.PUT("/chat/channels/:channel/users/:user")
		v2.DELETE("/chat/channels/:channel/users/:user")
		v2.PUT("/chat/channels/:channel/mark-as-read/:message")

		// === Comments ===
		v2.GET("/comments")
		v2.POST("/comments")
		v2.GET("/comments/:comment")
		v2.PUT("/comments/:comment")
		v2.PATCH("/comments/:comment")
		v2.DELETE("/comments/:comment")
		v2.POST("/comments/:comment/vote")
		v2.DELETE("/comments/:comment/vote")

		// === Notifications ===
		v2.GET("/notifications")
		v2.POST("/notifications/mark-read")

		// === Misc ===
		v2.POST("/reports")
		v2.GET("/changelog")
		v2.GET("/changelog/{changelog}")
	}
}
