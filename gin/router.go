package main

import (
	"fmt"
	"gin/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var urlHandller gin.HandlerFunc = func(c *gin.Context) {
	// name := c.Query("name") //if no query passed, then return empty string
	name := c.Query("name")
	// name := c.DefaultQuery("name", "default_name")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

var apiHandller gin.HandlerFunc = func(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	//截取/
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

var formHandler gin.HandlerFunc = func(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password ")
	c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, types: %s", username, password, types))
}

var uploader gin.HandlerFunc = func(c *gin.Context) {
	//fileKey is the key that passed in form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, fmt.Sprint("upload failed", err))
	}
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}

var uploaderV2 gin.HandlerFunc = func(c *gin.Context) {
	_, headers, err := c.Request.FormFile("file")
	if err != nil {
		c.String(500, fmt.Sprint("upload failed", err))
	}

	contentType := headers.Header.Get("Content-Type")
	fmt.Println(contentType)
	// if contentType != "text/plain" {
	allowedTypes := []string{"application/gzip", "image/jpeg", "text/plain"}

	// check uploaded file type is allowed or not
	if !utils.Contains(allowedTypes, contentType) {
		fmt.Printf("content type %v is not allowed", contentType)
		return
	}

	// check uploaded file size
	fmt.Println("----file size---", headers.Size)
	if headers.Size > 1024*1024*2 {
		fmt.Println("file size exceed limit, should below 2MB")
		return
	}

	timer := time.Now().Format("15:04:05")
	c.SaveUploadedFile(headers, timer+headers.Filename)
	c.String(http.StatusOK, headers.Filename)
}

func main() {
	r := gin.Default()
	// 20 times (8 times 2 ), 2^10 is 1024, so this means 8MB
	maxSize := 8 << 20

	// when uploading files, limit the max memory usage to 8MB (not the max file size)
	r.MaxMultipartMemory = int64(maxSize)

	//full url: http://localhost:8000/di/doSth
	r.GET("/:name/*action", apiHandller)
	r.GET("/user", urlHandller)
	r.POST("/form", formHandler)
	r.POST("/uploader", uploader)
	r.POST("/uploaderV2", uploaderV2)
	r.Run()
	// r.Run(":8000")
}
