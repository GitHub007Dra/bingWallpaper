package handle

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) *gin.Engine {
	common := r.Group("")
	{
		common.GET("/healthCheck", HealthCheck)
		common.GET("/bingWp", bingWallpaper)
	}

	return r
}
