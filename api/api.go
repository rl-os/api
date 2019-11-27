package api

import (
	v2_0 "github.com/deissh/osu-api-server/api/v2"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v2_0.ApplyRoutes(api)
	}
}
