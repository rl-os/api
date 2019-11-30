package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/deissh/osu-api-server/pkg/oauth/token"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	oauth := r.Group("/oauth")
	{
		oauth.POST("/token")
	}
}
