package router

import (
	"gin/controllers"
	"net/http"

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
		v1.GET("/session", controllers.SessionDemo)

		// redirect request
		v1.GET("/baidu", func(c *gin.Context) {
			// http.StatusMovedPermanently: 301 Moved Permanently
			c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		})
	}
}
