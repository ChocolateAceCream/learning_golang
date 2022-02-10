package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMiddleware() []gin.HandlerFunc {
	m := []gin.HandlerFunc{
		MiddlewareHandler,
		// SessionHandler,
	}
	return m
}

func MiddlewareHandler(c *gin.Context) {
	t := time.Now()
	fmt.Println("Start middleWare")

	// write key-value pair into context, so in downstream you can retrive value by c.Get(key)
	c.Set("timestamp", time.Now())
	status := c.Writer.Status()

	// Next should be used only inside middleware. It executes the pending handlers in the chain inside the calling handler.
	c.Next()
	fmt.Println("middleWare done", status)
	t2 := time.Since(t)
	fmt.Println("time used: ", t2)
}
