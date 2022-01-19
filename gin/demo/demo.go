package demo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func demo() {
	// 1.创建路由
	r := gin.Default()
	fmt.Println("----", 8<<20)
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
			"count":   123,
		})
	})

	r.Run() // default listen and serve on 0.0.0.0:8080
}
