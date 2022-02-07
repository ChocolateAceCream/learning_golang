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

	//init router
	router.SetupRouter(r)

	// init redis session store
	// services.SetupSession(r)
	return r
}

func main() {

	r := Init()

	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
	r.Run(":3000")
}
