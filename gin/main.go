package main

import (
	"fmt"
	"gin/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	//setup log
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// make sure gin.Default() is executed after init logger
	r := gin.Default()
	return r
}

func main() {

	r := Init()
	router.SetupRouter(r)

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	// r.Run(":8000")
}
