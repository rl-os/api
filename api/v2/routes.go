package v2

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": time.Now().Unix(),
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v2")
	{
		v1.GET("/ping", ping)
	}
}
