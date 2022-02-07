package services

import (
	"context"
	"gin/config"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var once = sync.Once{}
var currentClient *redis.Client

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
	newRedisClient.Set(ctx, "age1", 12333, 30*time.Second)
	currentClient = newRedisClient
	// currentClient.Set(ctx, "gbg", 123, 30*time.Second)
}

func RedisClient() *redis.Client {
	if currentClient == nil {
		once.Do(func() { newRedisClient() })
	}
	return currentClient
}
