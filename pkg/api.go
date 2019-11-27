package pkg

import (
	"github.com/deissh/osu-api-server/pkg/v2"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v2.ApplyRoutes(api)
	}
}
