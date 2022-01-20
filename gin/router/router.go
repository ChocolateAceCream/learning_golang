package router

import (
	"fmt"
	"gin/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func urlHandller(c *gin.Context) {
	// name := c.Query("name") //if no query passed, then return empty string
	name := c.Query("name")
	// name := c.DefaultQuery("name", "default_name")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

func apiHandller(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	//截取/
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

func formHandler(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password ")
	c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, types: %s", username, password, types))
}

// *gin.Engine is the thing that can be passed around
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 20 times (8 times 2 ), 2^10 is 1024, so this means 8MB
	maxSize := 8 << 20

	// when uploading files, limit the max memory usage to 8MB (not the max file size)
	r.MaxMultipartMemory = int64(maxSize)

	RouteLoader(r)
	//create route group to store similar routes
	v2 := r.Group("/v2")
	{
		v2.POST("/form", formHandler)
		v2.POST("/uploader", controllers.FileUploader)
		v2.POST("/uploaderV2", controllers.FileUploaderV2)
		v2.POST("/multiFilesUploader", controllers.MultiFilesUploader)
	}
	return r
}
