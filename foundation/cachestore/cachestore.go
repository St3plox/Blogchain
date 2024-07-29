package cachestore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cacheable is an interface that all cacheable entities should implement.
type Cacheable interface {
	CacheKey() string
	CacheExpiration() time.Duration
}

// CacheStorer is the interface for cache storage.
type CacheStorer interface {
	Set(ctx context.Context, value Cacheable) error
	Get(ctx context.Context, key string, target Cacheable) error
	Delete(ctx context.Context, key string) error
}

// RedisClient is a Redis implementation of CacheStorer.
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient creates a new RedisClient with the given redis.Client.
func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{Client: client}
}

// Set stores a Cacheable entity in Redis.
func (rc *RedisClient) Set(ctx context.Context, value Cacheable) error {
	key := value.CacheKey()
	expiration := value.CacheExpiration()

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	if err := rc.Client.Set(ctx, key, data, expiration).Err(); err != nil {
		return fmt.Errorf("failed to set value in Redis: %w", err)
	}

	return nil
}

// Get retrieves a Cacheable entity from Redis.
func (rc *RedisClient) Get(ctx context.Context, key string, target Cacheable) error {

	data, err := rc.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return err
	} else if err != nil {
		return fmt.Errorf("failed to get value from Redis: %w", err)
	}

	if err := json.Unmarshal([]byte(data), &target); err != nil {
		return fmt.Errorf("failed to unmarshal value: %w", err)
	}

	return nil
}

// Delete removes a value from Redis by key.
func (rc *RedisClient) Delete(ctx context.Context, key string) error {
	if err := rc.Client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete value from Redis: %w", err)
	}
	return nil
}
