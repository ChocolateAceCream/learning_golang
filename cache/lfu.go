/*
* @fileName lru.go
* @author Di Sheng
* @date 2025/04/24 09:25:12
* @description LFU Cache implementation in Go with redis
* TODO: By default when score the same, ZSet will evict based on ASCII/lexical order ("a" before "b"), now I changed to random keys from same freq set
 */
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

// T is the value type, should be marshable to JSON
type LFUCache[T any] struct {
	capacity    int
	store       *redis.Client
	dataKey     string
	freqKey     string
	totalCounts int
}

func orDefault[T comparable](value, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}

func ToRedisStorable(value any) (string, error) {
	// Just try marshaling everything
	jsonData, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("value cannot be stored in Redis: %v", err)
	}
	return string(jsonData), nil
}

func NewLFU[T any](options LFUCache[T]) *LFUCache[T] {
	return &LFUCache[T]{
		capacity: orDefault(options.capacity, 100),
		store:    options.store,
		dataKey:  orDefault(options.dataKey, "data"),
		freqKey:  orDefault(options.freqKey, "freq"),
	}
}

func (l *LFUCache[T]) Set(c context.Context, key string, value any) (err error) {
	val, err := ToRedisStorable(value)
	if err != nil {
		return
	}
	// check if key exists
	exists, err := l.store.HExists(c, l.dataKey, key).Result()
	if err != nil {
		return
	}
	if !exists {
		// if not, check if capacity is reached
		l.totalCounts++
		if l.totalCounts > l.capacity {
			// evict LFU key
			evictKey, err := l.store.ZRangeWithScores(c, l.freqKey, 0, 0).Result()

			if err != nil || len(evictKey) == 0 {
				return err
			}
			minFreq := evictKey[0].Score
			// instead of just delete the first key, we delete a random key with the same score
			targets, err := l.store.ZRangeByScore(c, l.freqKey, &redis.ZRangeBy{
				Min: fmt.Sprintf("%f", minFreq),
				Max: fmt.Sprintf("%f", minFreq),
			}).Result()
			if err != nil || len(targets) == 0 {
				return err
			}
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			key := targets[r.Intn(len(targets))]
			fmt.Println("--------evictKey---------", key)
			l.store.HDel(c, l.dataKey, key)
			l.store.ZRem(c, l.freqKey, key)
			l.totalCounts--
		}
	}
	// set the value and increment frequency
	err = l.store.HSet(c, l.dataKey, key, val).Err()
	if err != nil {
		return
	}
	l.store.ZIncrBy(c, l.freqKey, 1, key)

	return
}

func (l *LFUCache[T]) Get(c context.Context, key string) (value T, err error) {
	// check if key exists
	exists, err := l.store.HExists(c, l.dataKey, key).Result()
	if err != nil {
		return
	}
	if !exists {
		return value, fmt.Errorf("key %s not found", key)
	}
	// get the value and increment frequency
	val, err := l.store.HGet(c, l.dataKey, key).Result()
	if err != nil {
		return
	}
	l.store.ZIncrBy(c, l.freqKey, 1, key)
	// unmarshal the value
	err = json.Unmarshal([]byte(val), &value)
	if err != nil {
		return
	}
	return value, nil
}
