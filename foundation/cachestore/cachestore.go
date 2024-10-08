// Package cachestore provides a caching mechanism using Redis as the storage backend.
//
// This package defines interfaces and implementations for storing, retrieving, and deleting cacheable entities in Redis.
//
// The core components of this package include:
//
// - **Cacheable Interface**: An interface that should be implemented by all entities that are intended to be cached. It provides methods for getting the cache key and expiration duration.
//
// - **CacheStore Interface**: An interface for cache storage operations. It includes methods for setting, getting, and deleting cacheable entities.
//
// - **RedisClient**: A concrete implementation of the CacheStore interface using Redis. It provides methods to interact with Redis to store, retrieve, and delete cacheable entities.
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

// CacheStore is the interface for cache storage.
type CacheStore interface {
	Set(ctx context.Context, value Cacheable) error
	Get(ctx context.Context, key string, target Cacheable) error
	Delete(ctx context.Context, key string) error
}

// RedisClient is a Redis implementation of CacheStore.
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
