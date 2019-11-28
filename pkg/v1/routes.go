package v1

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I'm alive!",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		// Auth-free API endpoints (public data)
		v1.GET("/ping", ping)

		v1.GET("/surprise_me")
		v1.GET("/users")
		v1.GET("/users/whatid")
		v1.GET("/users/full")
		v1.GET("/users/achievements")
		v1.GET("/users/userpage")
		v1.GET("/users/lookup")
		v1.GET("/users/scores/best")
		v1.GET("/users/scores/recent")
		v1.GET("/badges")
		v1.GET("/badges/members")
		v1.GET("/beatmaps")
		v1.GET("/leaderboard")
		v1.GET("/scores")
		v1.GET("/beatmaps/rank_requests/status")

		// original api
		v1.GET("/get_user")
		v1.GET("/get_match")
		v1.GET("/get_user_recent")
		v1.GET("/get_user_best")
		v1.GET("/get_scores")
		v1.GET("/get_beatmaps")

		// Auth
		v1.POST("/tokens/self/delete")

		// Tokens
		v1.GET("/tokens")
		v1.GET("/users/self")
		v1.GET("/tokens/self")

		// ReadConfidential privilege required
		v1.GET("/friends")
		v1.GET("/friends/with")
		v1.GET("/users/self/donor_info")
		v1.GET("/users/self/favourite_mode")
		v1.GET("/users/self/settings")

		// Write privilege required
		v1.POST("/friends/add")
		v1.POST("/friends/del")
		v1.POST("/users/self/settings")
		v1.POST("/users/self/userpage")
		v1.POST("/beatmaps/rank_requests")
	}
}
