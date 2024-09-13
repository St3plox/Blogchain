package like_test

import (
	"context"
	"errors"
	"testing"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/broker"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockStorer struct {
	mock.Mock
}

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) ProduceLikesEvents(likesEvents []broker.LikeEvent) error {
	args := m.Called(likesEvents)
	return args.Error(0)
}

func (m *MockStorer) Create(ctx context.Context, newLike like.Like) (like.Like, error) {
	args := m.Called(ctx, newLike)
	return args.Get(0).(like.Like), args.Error(1)
}

func (m *MockStorer) QueryByID(ctx context.Context, likeID string) (like.Like, error) {
	args := m.Called(ctx, likeID)
	return args.Get(0).(like.Like), args.Error(1)
}

func (m *MockStorer) QueryAllByUserID(ctx context.Context, userID string) ([]like.Like, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]like.Like), args.Error(1)
}

func (m *MockStorer) QueryAllByPostID(ctx context.Context, postID string) ([]like.Like, error) {
	args := m.Called(ctx, postID)
	return args.Get(0).([]like.Like), args.Error(1)
}

func (m *MockStorer) Update(ctx context.Context, updatedLike like.Like) (like.Like, error) {
	args := m.Called(ctx, updatedLike)
	return args.Get(0).(like.Like), args.Error(1)
}

func (m *MockStorer) DeleteByID(ctx context.Context, likeID string) error {
	args := m.Called(ctx, likeID)
	return args.Error(0)
}

func TestCore_Create(t *testing.T) {
	ctx := context.Background()

	mockCache := new(cachestore.MockCacheStore)
	mockStorer := new(MockStorer)
	mockUserStore := new(user.MockStorer)
	mockProducer := new(MockProducer)

	core := like.NewCore(mockCache, mockStorer, mockUserStore, mockProducer)

	// Mocked claims
	userID := primitive.NewObjectID()
	mockUser := user.User{
		ID: userID,
	}
	newLike := like.Like{
		PostID:     1,
		IsPositive: true,
	}
	expectedLike := newLike
	expectedLike.UserID = userID
	expectedLike.ID = primitive.NewObjectID()

	// Setup Mocks
	mockUserStore.On("QueryByID", ctx, mock.Anything).Return(mockUser, nil)
	mockStorer.On("Create", ctx, mock.Anything).Return(expectedLike, nil)
	mockCache.On("Set", ctx, mock.Anything).Return(nil)
	mockProducer.On("ProduceLikesEvents", mock.Anything).Return(nil)

	// Act
	createdLike, err := core.Create(ctx, newLike)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedLike, createdLike)
	mockUserStore.AssertCalled(t, "QueryByID", ctx, mock.Anything)
	mockStorer.AssertCalled(t, "Create", ctx, mock.Anything)
	mockCache.AssertCalled(t, "Set", ctx, expectedLike)
	mockProducer.AssertCalled(t, "ProduceLikesEvents", []broker.LikeEvent{
        {
            UserID:     expectedLike.UserID.Hex(),
            PostID:     expectedLike.PostID,
            IsPositive: expectedLike.IsPositive,
            UserEmail:  mockUser.Email,
        },
    })
}
func TestCore_QueryByID_CacheHit(t *testing.T) {
	ctx := context.Background()

	mockCache := new(cachestore.MockCacheStore)
	mockStorer := new(MockStorer)

	core := like.NewCore(mockCache, mockStorer, nil, nil)

	likeID := primitive.NewObjectID().Hex()
	expectedLike := like.Like{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PostID:     1,
		IsPositive: true,
	}

	var cachedLike like.Like
	mockCache.On("Get", ctx, like.IdToCacheKey(likeID), &cachedLike).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*like.Like)
		*arg = expectedLike
	}).Return(nil)

	// Act
	resultLike, err := core.QueryByID(ctx, likeID)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedLike, resultLike)
	mockCache.AssertExpectations(t)
	mockStorer.AssertExpectations(t)
}

func TestCore_QueryByID_CacheMiss(t *testing.T) {
	ctx := context.Background()

	mockCache := new(cachestore.MockCacheStore)
	mockStorer := new(MockStorer)

	core := like.NewCore(mockCache, mockStorer, nil, nil)

	likeID := primitive.NewObjectID().Hex()
	expectedLike := like.Like{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PostID:     1,
		IsPositive: true,
	}

	var cachedLike like.Like
	mockCache.On("Get", ctx, like.IdToCacheKey(likeID), &cachedLike).Return(errors.New("cache miss"))
	mockStorer.On("QueryByID", ctx, likeID).Return(expectedLike, nil)
	mockCache.On("Set", ctx, expectedLike).Return(nil)

	// Act
	resultLike, err := core.QueryByID(ctx, likeID)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedLike, resultLike)
	mockCache.AssertExpectations(t)
	mockStorer.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func TestCore_Update(t *testing.T) {
	ctx := context.Background()

	mockCache := new(cachestore.MockCacheStore)
	mockStorer := new(MockStorer)

	core := like.NewCore(mockCache, mockStorer, nil, nil)

	updatedLike := like.Like{
		ID:         primitive.NewObjectID(),
		UserID:     primitive.NewObjectID(),
		PostID:     1,
		IsPositive: false,
	}

	// Setup Mocks
	mockStorer.On("Update", ctx, updatedLike).Return(updatedLike, nil)
	mockCache.On("Set", ctx, updatedLike).Return(nil)

	// Act
	resultLike, err := core.Update(ctx, updatedLike)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, updatedLike, resultLike)
	mockStorer.AssertCalled(t, "Update", ctx, updatedLike)
	mockCache.AssertCalled(t, "Set", ctx, updatedLike)
}

func TestCore_DeleteByID(t *testing.T) {
	ctx := context.Background()

	mockCache := new(cachestore.MockCacheStore)
	mockStorer := new(MockStorer)

	core := like.NewCore(mockCache, mockStorer, nil, nil)

	likeID := primitive.NewObjectID().Hex()

	// Setup Mocks
	mockCache.On("Delete", ctx, like.IdToCacheKey(likeID)).Return(nil)
	mockStorer.On("DeleteByID", ctx, likeID).Return(nil)

	// Act
	err := core.DeleteByID(ctx, likeID)

	// Assert
	require.NoError(t, err)
	mockCache.AssertCalled(t, "Delete", ctx, like.IdToCacheKey(likeID))
	mockStorer.AssertCalled(t, "DeleteByID", ctx, likeID)
}
