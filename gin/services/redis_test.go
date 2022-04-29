package services

import (
	"fmt"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestRedis(t *testing.T) {
	fmt.Println("----testing----")
	wait := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wait.Add(1)
		go func(index int) {

			defer wait.Done()
			time.Sleep(time.Duration(index) * time.Second)
			fmt.Println("----index-----", index)
			lock := GetLock("test", 1*time.Second)
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			defer lock.Release(c)
			if lock.AcquireLock(c) != "" {
				fmt.Println("拿锁成功:", index)
				time.Sleep(4 * time.Second)
			} else {
				fmt.Println("拿锁失败:", index)

			}
		}(i)
	}
	wait.Wait()
	fmt.Println("finish")
}
