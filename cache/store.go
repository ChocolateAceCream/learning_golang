package cache

import (
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type Option func(*RedisConfig)

func WithAddr(addr string) Option {
	return func(c *RedisConfig) {
		c.Addr = addr
	}
}

func WithPassword(password string) Option {
	return func(c *RedisConfig) {
		c.Password = password
	}
}

func WithDB(db int) Option {
	return func(c *RedisConfig) {
		c.DB = db
	}
}

func NewRedisClient(opts ...Option) *redis.Client {
	// Default config
	config := &RedisConfig{
		Addr: "localhost:6379",
		DB:   2, // Or another default
	}

	// Apply options
	for _, opt := range opts {
		opt(config)
	}

	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
}
