package services

import (
	"context"
	"fmt"
	"gin/config"
	"gin/utils"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var once = sync.Once{}
var currentClient *redis.Client

type Lock struct {
	redis *redis.Client
}

func GetLock() *Lock {
	return &Lock{
		redis: GetRedisClient(),
	}
}

func newRedisClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//read env variables from config file
	config := config.GetConfig()
	redisHost := config.Redis.Host
	redisPassword := config.Redis.Password

	// alternatively, read env var from system env
	// redisHost := os.Getenv("REDIS_HOST")
	// redisPassword := os.Getenv("REDIS_PW")
	conf := &redis.Options{
		Addr:     redisHost,
		DB:       1,
		Password: redisPassword,
	}
	newRedisClient := redis.NewClient(conf)
	resp := newRedisClient.Ping(ctx)
	if resp.Err() != nil {
		panic(resp.Err())
	}
	currentClient = newRedisClient

	// currentClient.Set(ctx, "gbg", 123, 30*time.Second)
}

// using sync.once to only init redis client once
func GetRedisClient() *redis.Client {
	if currentClient == nil {
		once.Do(func() { newRedisClient() })
	}
	return currentClient
}

// return lockId which used to identify locks
// this lock is blocked lock.
/* usage example:
lock := GetLock()
// lockId := lock.AcquireLock(c, lockName, 5*time.Second) // attemp to acquire a lock which has expiration time of 5 seconds
// defer lock.Release(c, lockName, lockId)
*/
func (lock *Lock) AcquireLock(c *gin.Context, locakname string, expiration time.Duration) string {
	lockId := utils.RandToken(10)
	key := "lock:" + locakname

	// set tick interval to 100ns, which try to acquire lock every 100ns
	tick := time.NewTicker(time.Nanosecond * 100)

	// set time out to 10 second
	timer := time.NewTimer(10 * time.Second)

	for {
		select {

		// time out
		case <-timer.C:
			timer.Stop()
			return ""

		case <-tick.C:
			setNxCmd := lock.redis.SetNX(c, key, lockId, expiration)
			if ok, _ := setNxCmd.Result(); ok {
				return lockId
			}
		}
	}
}

func (lock *Lock) Release(c *gin.Context, locakname, lockId string) bool {
	key := "lock:" + locakname

	// alternatively, we can use redis pipeline
	// timer := time.NewTimer(50 * time.Second)
	// txf := func(tx *redis.Tx) error {
	// 	getCmd := tx.Get(c, key)
	// 	fn := func(pipe redis.Pipeliner) error {
	// 		if getCmd.Val() == lockId {
	// 			pipe.Del(c, key)
	// 		}
	// 		return nil
	// 	}
	// 	_, err := tx.TxPipelined(c, fn)
	// 	return err
	// }

	// for {
	// 	select {
	// 	case <-timer.C:
	// 		timer.Stop()
	// 		return false
	// 	default:
	// 		err := lock.redis.Watch(c, txf, key)
	// 		if err == nil {
	// 			return true
	// 		} else if err == redis.TxFailedErr {
	// 			// something wrong, we either lost the key or an unexpected thing happened, just try again
	// 			continue
	// 		} else {
	// 			return false
	// 		}
	// 	}
	// }

	// or we can use lua script to implement atomic operation
	const luaScript = `
	if redis.call('get', KEYS[1])==ARGV[1] then
		return redis.call('del', KEYS[1])
	else
		return 0
	end
	`

	script := redis.NewScript(luaScript)
	result, err := script.Run(c, lock.redis, []string{key}, lockId).Result()
	fmt.Println("-----lock release-----", err)
	if err == nil {
		return result == 1 // 1: delete successful
	} else {
		return false
	}
}
