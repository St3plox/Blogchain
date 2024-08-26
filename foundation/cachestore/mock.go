package cachestore

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockCacheStore mocks the CacheStorer interface.
type MockCacheStore struct {
	mock.Mock
}

func (m *MockCacheStore) Set(ctx context.Context, val Cacheable) error {
	args := m.Called(ctx, val)
	return args.Error(0)
}

func (m *MockCacheStore) Get(ctx context.Context, key string, val Cacheable) error {
	args := m.Called(ctx, key, val)
	return args.Error(0)
}

func (m *MockCacheStore) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}
