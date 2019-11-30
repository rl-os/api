package v1

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
	v1 := r.Group("/v1")
	{
		// Auth-free API endpoints (public data)
		v1.GET("/empty", empty)

		v1.GET("/surprise_me", empty)
		v1.GET("/users", empty)
		v1.GET("/users/whatid", empty)
		v1.GET("/users/full", empty)
		v1.GET("/users/achievements", empty)
		v1.GET("/users/userpage", empty)
		v1.GET("/users/lookup", empty)
		v1.GET("/users/scores/best", empty)
		v1.GET("/users/scores/recent", empty)
		v1.GET("/badges", empty)
		v1.GET("/badges/members", empty)
		v1.GET("/beatmaps", empty)
		v1.GET("/leaderboard", empty)
		v1.GET("/scores", empty)
		v1.GET("/beatmaps/rank_requests/status", empty)

		// original api
		v1.GET("/get_user", empty)
		v1.GET("/get_match", empty)
		v1.GET("/get_user_recent", empty)
		v1.GET("/get_user_best", empty)
		v1.GET("/get_scores", empty)
		v1.GET("/get_beatmaps", empty)

		// Auth
		v1.POST("/tokens/self/delete", empty)

		// Tokens
		v1.GET("/tokens", empty)
		v1.GET("/users/self", empty)
		v1.GET("/tokens/self", empty)

		// ReadConfidential privilege required
		v1.GET("/friends", empty)
		v1.GET("/friends/with", empty)
		v1.GET("/users/self/donor_info", empty)
		v1.GET("/users/self/favourite_mode", empty)
		v1.GET("/users/self/settings", empty)

		// Write privilege required
		v1.POST("/friends/add", empty)
		v1.POST("/friends/del", empty)
		v1.POST("/users/self/settings", empty)
		v1.POST("/users/self/userpage", empty)
		v1.POST("/beatmaps/rank_requests", empty)
	}
}
