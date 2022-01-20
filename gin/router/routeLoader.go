package router

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func RouteLoader(r *gin.Engine) {
	v1 := r.Group("/v1")
	// {} is standard format for route group
	{
		v1.GET("/:name/*action", apiHandller)
		v1.GET("/user", urlHandller)
		v1.POST("/login/", controllers.Login)
		v1.GET("/info/:id", controllers.GetInfo)
	}
}
