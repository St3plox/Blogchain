package media_test

import (
	"context"
	"net/mail"
	"testing"

	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockStorer struct {
	mock.Mock
}

type mockUserStorer struct {
	mock.Mock
}

type mockCacheStore struct {
	mock.Mock
}

func (m *mockStorer) Create(ctx context.Context, newMedia media.Media) (media.Media, error) {
	args := m.Called(ctx, newMedia)
	return args.Get(0).(media.Media), args.Error(1)
}

func (m *mockStorer) CreateMultiple(ctx context.Context, newMedia []media.Media) ([]media.Media, error) {
	args := m.Called(ctx, newMedia)
	return args.Get(0).([]media.Media), args.Error(1)
}

func (m *mockStorer) Update(ctx context.Context, newMedia media.Media) (media.Media, error) {
	args := m.Called(ctx, newMedia)
	return args.Get(0).(media.Media), args.Error(1)
}

func (m *mockStorer) Delete(ctx context.Context, media media.Media) error {
	args := m.Called(ctx, media)
	return args.Error(0)
}

func (m *mockStorer) DeleteByID(ctx context.Context, mediaID string) error {
	args := m.Called(ctx, mediaID)
	return args.Error(0)
}

func (m *mockStorer) QueryByID(ctx context.Context, mediaID string) (media.Media, error) {
	args := m.Called(ctx, mediaID)
	return args.Get(0).(media.Media), args.Error(1)
}

func (m *mockStorer) QueryByIDs(ctx context.Context, mediaIDs []string) ([]media.Media, error) {
	args := m.Called(ctx, mediaIDs)
	return args.Get(0).([]media.Media), args.Error(1)
}

type MockStorer struct {
	mock.Mock
}

func (m *MockStorer) Create(ctx context.Context, usr user.User) (user.User, error) {
	args := m.Called(ctx, usr)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockStorer) Delete(ctx context.Context, usr user.User) error {
	args := m.Called(ctx, usr)
	return args.Error(0)
}

func (m *MockStorer) Query(ctx context.Context, filter user.QueryFilter, orderBy string, pageNumber int, rowsPerPage int) ([]user.User, error) {
	args := m.Called(ctx, filter, orderBy, pageNumber, rowsPerPage)
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *MockStorer) Count(ctx context.Context, filter user.QueryFilter) (int, error) {
	args := m.Called(ctx, filter)
	return args.Int(0), args.Error(1)
}

func (m *MockStorer) QueryByID(ctx context.Context, userID string) (user.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(user.User), args.Error(1)
}

func (m *MockStorer) QueryByIDs(ctx context.Context, userIDs []string) ([]user.User, error) {
	args := m.Called(ctx, userIDs)
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *MockStorer) QueryByEmail(ctx context.Context, email mail.Address) (user.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(user.User), args.Error(1)
}

func TestCore_Create(t *testing.T) {
	ctx := context.Background()
	mockStorer := new(mockStorer)
	mockCache := new(cachestore.MockCacheStore)
	mockUserStorer := new(MockStorer)

	core := media.NewCore(mockStorer, mockCache, mockUserStorer)

	userID := primitive.NewObjectID()

	// Set up user
	claims := auth.Claims{RegisteredClaims: jwt.RegisteredClaims{
		Subject: userID.Hex(),
	}}
	ctx = auth.SetClaims(ctx, claims)

	user := user.User{ID: userID}
	mockUserStorer.On("QueryByID", ctx, claims.Subject).Return(user, nil)

	// Set up media creation
	newMedia := media.NewMedia{
		Filename:  "file.jpg",
		Length:    12345,
		FileBytes: []byte{1, 2, 3},
	}

	expectedMedia := media.Media{
		ID:        primitive.NewObjectID(),
		Filename:  newMedia.Filename,
		Length:    newMedia.Length,
		FileBytes: newMedia.FileBytes,
		OwnerID:   user.ID,
	}

	mockStorer.On("Create", ctx, mock.AnythingOfType("media.Media")).Return(expectedMedia, nil)

	createdMedia, err := core.Create(ctx, newMedia)
	assert.NoError(t, err)
	assert.Equal(t, newMedia.Filename, createdMedia.Filename)

	mockUserStorer.AssertExpectations(t)
	mockStorer.AssertExpectations(t)
}

func TestCore_DeleteByID_OwnerCheck(t *testing.T) {
	ctx := context.Background()
	mockStorer := new(mockStorer)
	mockCache := new(cachestore.MockCacheStore)
	mockUserStorer := new(MockStorer)

	core := media.NewCore(mockStorer, mockCache, mockUserStorer)

	userID := primitive.NewObjectID()

	// Set up user
	claims := auth.Claims{RegisteredClaims: jwt.RegisteredClaims{
		Subject: userID.Hex(),
	}}
	ctx = auth.SetClaims(ctx, claims)

	user := user.User{ID: userID}
	mockUserStorer.On("QueryByID", ctx, claims.Subject).Return(user, nil)

	// Media not owned by the user
	testMedia := media.Media{
		ID:      primitive.NewObjectID(),
		OwnerID: primitive.NewObjectID(), // Different owner
	}
	mockStorer.On("QueryByID", ctx, testMedia.ID.Hex()).Return(testMedia, nil)

	err := core.DeleteByID(ctx, testMedia.ID.Hex())
	assert.ErrorIs(t, err, media.ErrUserNotOwner)
}

func TestCore_QueryByID_CacheHit(t *testing.T) {
    ctx := context.Background()
    mockStorer := new(mockStorer)
    mockCache := new(cachestore.MockCacheStore) // Correctly initializing the cache mock

    core := media.NewCore(mockStorer, mockCache, nil)

    mediaID := "some-media-id"
    cachedMedia := media.Media{ID: primitive.NewObjectID(), Filename: "cached_file.jpg"}

    // Simulate cache hit by setting the cache to return the correct data
    mockCache.On("Get", ctx, media.IdToCacheKey(mediaID), mock.AnythingOfType("*media.Media")).Run(func(args mock.Arguments) {
        dest := args.Get(2).(*media.Media)
        *dest = cachedMedia // Assign the cached media to the result
    }).Return(nil)

    // Call the core's QueryByID method
    result, err := core.QueryByID(ctx, mediaID)

    // Assertions
    assert.NoError(t, err)
    assert.Equal(t, cachedMedia.Filename, result.Filename)

    // Ensure that the cache was checked, and the storer was never called
    mockCache.AssertExpectations(t)
    mockStorer.AssertNotCalled(t, "QueryByID") // Ensure the storer is never called
}


func TestCore_QueryByID_CacheMiss_DBHit(t *testing.T) {
	ctx := context.Background()
	mockStorer := new(mockStorer)
	mockCache := new(cachestore.MockCacheStore)

	core := media.NewCore(mockStorer, mockCache, nil)

	mediaID := "some-media-id"
	dbMedia := media.Media{ID: primitive.NewObjectID(), Filename: "db_file.jpg"}

	// Simulate cache miss and DB hit
	mockCache.On("Get", ctx, media.IdToCacheKey(mediaID), mock.AnythingOfType("*media.Media")).Return(redis.Nil)
	mockStorer.On("QueryByID", ctx, mediaID).Return(dbMedia, nil)
	mockCache.On("Set", ctx, dbMedia).Return(nil)

	result, err := core.QueryByID(ctx, mediaID)
	assert.NoError(t, err)
	assert.Equal(t, dbMedia.Filename, result.Filename)

	mockCache.AssertExpectations(t)
	mockStorer.AssertExpectations(t)
}

func TestCore_QueryByIDs(t *testing.T) {
	ctx := context.Background()
	mockStorer := new(mockStorer)
	mockCache := new(cachestore.MockCacheStore)

	core := media.NewCore(mockStorer, mockCache, nil)

	mediaIDs := []string{"id1", "id2"}
	cachedMedia := media.Media{ID: primitive.NewObjectID(), Filename: "cached_file.jpg"}
	dbMedia := media.Media{ID: primitive.NewObjectID(), Filename: "db_file.jpg"}

	// Cache miss for id1, cache hit for id2
	mockCache.On("Get", ctx, media.IdToCacheKey("id1"), mock.AnythingOfType("*media.Media")).Return(redis.Nil)
	mockCache.On("Get", ctx, media.IdToCacheKey("id2"), mock.AnythingOfType("*media.Media")).Run(func(args mock.Arguments) {
		dest := args.Get(2).(*media.Media)
		*dest = cachedMedia
	}).Return(nil)

	// DB hit for id1
	mockStorer.On("QueryByIDs", ctx, []string{"id1"}).Return([]media.Media{dbMedia}, nil)
	mockCache.On("Set", ctx, dbMedia).Return(nil)

	result, err := core.QueryByIDs(ctx, mediaIDs)
	assert.NoError(t, err)
	assert.Len(t, result, 2)

	mockCache.AssertExpectations(t)
	mockStorer.AssertExpectations(t)
}
