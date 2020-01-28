package v2

import (
	"github.com/deissh/osu-api-server/pkg/middlewares/permission"
	"github.com/deissh/osu-api-server/pkg/v2/channels"
	"github.com/deissh/osu-api-server/pkg/v2/chats"
	"github.com/deissh/osu-api-server/pkg/v2/users"
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
		v2.GET("/me/", users.GetUserByToken, permission.MustLogin)
		v2.GET("/me/:mode", users.GetUserByToken, permission.MustLogin)
		v2.GET("/me/download-quota-check", empty, permission.MustLogin)

		// === Friends ===
		v2.GET("/friends", empty, permission.MustLogin)

		// === Users ===
		v2.GET("/users/:user/kudosu", empty)
		v2.GET("/users/:user/scores/:type", empty)
		v2.GET("/users/:user/beatmapsets/:type", empty)
		v2.GET("/users/:user/recent_activity", empty)
		v2.GET("/users/:user/:mode", users.GetUserByID)
		v2.GET("/users/:user", users.GetUserByID)

		// === Beatmaps ===
		v2.GET("/beatmaps/lookup", empty)
		v2.GET("/beatmaps/:beatmap", empty)
		v2.GET("/beatmaps/:beatmap/scores", empty)
		v2.GET("/beatmapsets/lookup", empty)
		v2.GET("/beatmapsets/search/:filters", empty)
		v2.GET("/beatmapsets/:beatmapset", empty)
		v2.GET("/beatmapsets/:beatmapset/download", empty, permission.IsSupporter)
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
		v2Chat := v2.Group("/chat")
		{
			v2Chat.POST("/new", empty)
			v2Chat.GET("/updates", empty)
			v2Chat.GET("/presence", empty) // ???
			v2Chat.GET("/channels", channels.GetAll)
			v2Chat.GET("/channels/joined", channels.GetJoinedAll)
			v2Chat.GET("/channels/:channel/messages", empty)
			v2Chat.POST("/channels/:channel/messages", chats.ChannelSendHandler)
			v2Chat.PUT("/channels/:channel/users/:user", empty)
			v2Chat.DELETE("/channels/:channel/users/:user", empty)
			v2Chat.PUT("/channels/:channel/mark-as-read/:message", empty)
		}

		// === Comments ===
		v2Comments := v2.Group("/comments")
		{
			v2Comments.GET("/", empty)
			v2Comments.POST("/", empty)
			v2Comments.GET("/:comment", empty)
			v2Comments.PUT("/:comment", empty)
			v2Comments.PATCH("/:comment", empty)
			v2Comments.DELETE("/:comment", empty)
			v2Comments.POST("/:comment/vote", empty)
			v2Comments.DELETE("/:comment/vote", empty)
		}

		// === Notifications ===
		v2.GET("/notifications", empty)
		v2.POST("/notifications/mark-read", empty)

		// === Misc ===
		v2.POST("/reports", empty)
		v2.GET("/changelog", empty)
		v2.GET("/changelog/:changelog", empty)
	}
}
