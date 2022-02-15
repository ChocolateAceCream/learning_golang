package controllers

import (
	"fmt"
	"gin/middleware"
	"gin/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginQuery struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"passwd" json:"password" uri:"pw" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var parsedQuery LoginQuery

	// request with body in JSON format
	//ShouldBindJSON() will bind json from request body into parsedQuery obj
	// if err := c.BindJSON(&parsedQuery); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
	// 	return
	// }

	// request with body in form format
	// Bind() will bind form query from request body into parsedQuery obj
	if err := c.Bind(&parsedQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}

	fmt.Println("----- username----", parsedQuery.User)
	fmt.Println("----- Password----", parsedQuery.Password)
	if parsedQuery.User != "admin" || parsedQuery.Password != "123qwe" {
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": "not authorized"})
		return
	}
	services.GetRedisClient().Set(c, "username", parsedQuery.User, 30*time.Second)

}

// GET localhost:3000/v1/info/1?username=admin&passwd=123qwe
func GetInfo(c *gin.Context) {
	var parsedQuery LoginQuery
	id := c.Param("id")
	fmt.Println("----- id----", id)
	if err := c.BindQuery(&parsedQuery); err != nil {
		fmt.Println("----- parsedQuery----", parsedQuery)
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
	}
	if parsedQuery.User != "admin" || parsedQuery.Password != "123qwe" {
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": "not authorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200", "from": "getinfo"})
}

func LockDemo(c *gin.Context) {
	ip := c.ClientIP()
	fmt.Println("------ip---", ip)
	opts := services.LockOptions{
		Attempts:        10,
		Ip:              ip,
		Duration:        10 * time.Second,
		EndPoint:        "LockDemo",
		BlockedDuration: 5 * time.Second,
	}
	rateLimiter, _ := services.RateLimiter(c, opts)

	c.Header("Content-Type", "application/json")
	if rateLimiter.Block {
		c.Writer.WriteHeader(http.StatusTooManyRequests)
		timeLeft := float32(rateLimiter.TimeLeft) / float32(time.Microsecond)
		c.Writer.Write([]byte(fmt.Sprintf("You reached the limit, come back after %vms", timeLeft)))
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	attemptsLeft := strconv.Itoa(rateLimiter.AttemptsLeft)
	c.Writer.Write([]byte("Attempts Left: " + attemptsLeft))
}

func SessionDemo(c *gin.Context) {
	// session set demo
	session := middleware.GetSession(c)
	err := session.Set("aaa", "aaaaa")
	if err != nil {
		fmt.Printf("err: ", err)
	}

	// session get bey key demo
	str, err := session.Get("1644304326")
	if err != nil {
		fmt.Printf("err: ", err)
	}

	// session remove by key demo
	err = session.Remove("aaa")
	if err != nil {
		fmt.Printf("err: ", err)
	}
	fmt.Println("---get session -----", str)

	for i := 0; i < 100; i++ {
		go func(index int) {
			session.Set(strconv.Itoa(index), index)
		}(i)
	}

	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
