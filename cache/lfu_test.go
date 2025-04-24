package cache

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func setupTestRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   9, // Use test DB to avoid clashing
	})
}

func cleanupRedisKeys(client *redis.Client, keys ...string) {
	for _, key := range keys {
		client.Del(context.Background(), key)
	}
}

func TestLFUCache_SetAndGet(t *testing.T) {
	client := setupTestRedisClient()
	cache := NewLFU[string](LFUCache[string]{
		capacity: 2,
		store:    client,
		dataKey:  "test:data",
		freqKey:  "test:freq",
	})

	ctx := context.Background()
	defer cleanupRedisKeys(client, cache.dataKey, cache.freqKey)

	err := cache.Set(ctx, "a", "alpha")
	assert.NoError(t, err)

	err = cache.Set(ctx, "b", "beta")
	assert.NoError(t, err)

	val, err := cache.Get(ctx, "a")
	assert.NoError(t, err)
	assert.Equal(t, "alpha", val)

	// Touch "a" again so it has higher frequency than "b"
	_, _ = cache.Get(ctx, "a")

	// Add a new item to trigger eviction
	err = cache.Set(ctx, "c", "gamma")
	assert.NoError(t, err)

	// Now "b" should be evicted (LFU)
	_, err = cache.Get(ctx, "b")
	assert.Error(t, err)

	// "a" and "c" should still be present
	v1, err := cache.Get(ctx, "a")
	assert.NoError(t, err)
	assert.Equal(t, "alpha", v1)

	v2, err := cache.Get(ctx, "c")
	assert.NoError(t, err)
	assert.Equal(t, "gamma", v2)
}

func TestLFUCache_GetNonExistentKey(t *testing.T) {
	client := setupTestRedisClient()
	cache := NewLFU[string](LFUCache[string]{
		capacity: 2,
		store:    client,
		dataKey:  "test:data",
		freqKey:  "test:freq",
	})

	ctx := context.Background()
	defer cleanupRedisKeys(client, cache.dataKey, cache.freqKey)

	_, err := cache.Get(ctx, "not-there")
	assert.Error(t, err)
}
