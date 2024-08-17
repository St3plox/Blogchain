package cachestore_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// TestCacheable is a simple implementation of the Cacheable interface for testing purposes
type TestCacheable struct {
	Key        string
	Expiration time.Duration
}

// CacheKey returns the cache key for the item
func (t *TestCacheable) CacheKey() string {
	return t.Key
}

// CacheExpiration returns the duration for which the item should be cached
func (t *TestCacheable) CacheExpiration() time.Duration {
	return t.Expiration
}

func TestRedisClient_Set(t *testing.T) {
	t.Run("test Set", func(t *testing.T) {

		const wantKey = "new-key"

		wantValue := TestCacheable{
			Key:        wantKey,
			Expiration: 5 * time.Minute,
		}

		// set up Miniredis
		mr := miniredis.RunT(t)

		// Set up the client
		rc := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})

		defer func() {
			rc.Close()
			mr.Close()
		}()

		cacheClient := &cachestore.RedisClient{
			Client: rc,
		}

		// Set the cacheable item in Redis
		err := cacheClient.Set(context.Background(), &wantValue)
		assert.Nil(t, err)

		// Verify the item was correctly set in Redis
		val, err := mr.Get(wantKey)
		assert.NoError(t, err)
		assert.NotEqual(t, "", val)

		var gotValue TestCacheable
		err = json.Unmarshal([]byte(val), &gotValue)
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)

		// Since miniredis does not automatically decrement TTL, manually fast-forward time to simulate expiration
		mr.FastForward(wantValue.Expiration)

		// Verify that the key has expired
		val, err = mr.Get(wantKey)
		assert.Equal(t, "", val)
	})
}

func TestRedisClient_Get(t *testing.T) {
	t.Run("test Get", func(t *testing.T) {

		const wantKey = "existing-key"

		wantValue := TestCacheable{
			Key:        wantKey,
			Expiration: 5 * time.Minute,
		}

		// set up Miniredis
		mr := miniredis.RunT(t)

		// Set up the client
		rc := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})

		defer func() {
			rc.Close()
			mr.Close()
		}()

		cacheClient := &cachestore.RedisClient{
			Client: rc,
		}

		// Set the cacheable item in Redis
		err := cacheClient.Set(context.Background(), &wantValue)
		assert.Nil(t, err)

		// Get the cacheable item from Redis
		var gotValue TestCacheable
		err = cacheClient.Get(context.Background(), wantKey, &gotValue)
		assert.Nil(t, err)
		assert.Equal(t, wantValue, gotValue)
	})
}

func TestRedisClient_Delete(t *testing.T) {
	t.Run("test Delete", func(t *testing.T) {

		const wantKey = "key-to-delete"

		wantValue := TestCacheable{
			Key:        wantKey,
			Expiration: 5 * time.Minute,
		}

		// set up Miniredis
		mr := miniredis.RunT(t)

		// Set up the client
		rc := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})

		defer func() {
			rc.Close()
			mr.Close()
		}()

		cacheClient := &cachestore.RedisClient{
			Client: rc,
		}

		// Set the cacheable item in Redis
		err := cacheClient.Set(context.Background(), &wantValue)
		assert.Nil(t, err)

		// Delete the cacheable item from Redis
		err = cacheClient.Delete(context.Background(), wantKey)
		assert.Nil(t, err)

		// Try to get the deleted item
		var gotValue TestCacheable
		err = cacheClient.Get(context.Background(), wantKey, &gotValue)
		assert.NotNil(t, err)
		assert.Equal(t, gotValue, TestCacheable{})
	})
}
