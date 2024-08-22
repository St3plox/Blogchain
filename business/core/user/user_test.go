package user_test

import (
	"context"
	"net/mail"
	"testing"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/foundation/blockchain"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// MockStorer mocks the Storer interface.
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

// MockCacheStorer mocks the CacheStorer interface.
type MockCacheStorer struct {
	mock.Mock
}

func (m *MockCacheStorer) Set(ctx context.Context, val cachestore.Cacheable) error {
	args := m.Called(ctx, val)
	return args.Error(0)
}

func (m *MockCacheStorer) Get(ctx context.Context, key string, val cachestore.Cacheable) error {
	args := m.Called(ctx, key, val)
	return args.Error(0)
}

func (m *MockCacheStorer) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func TestCore_QueryByIDFromCahe(t *testing.T) {
    // Setup
    mockCacheStorer := new(MockCacheStorer)
    mockStorer := new(MockStorer)
    core, _ := user.NewCore(mockStorer, nil, mockCacheStorer)

    expectedUser := user.User{
        ID:    primitive.NewObjectID(),
        Name:  "John Doe",
        Email: "john.doe@example.com",
	}
    

    // Mock the cache store Get method
    mockCacheStorer.On("Get", mock.Anything, "user:123", mock.AnythingOfType("*user.User")).
        Run(func(args mock.Arguments) {
            usr := args.Get(2).(*user.User)
            *usr = expectedUser
        }).
        Return(nil).
        Once()

    // Test
    usr, err := core.QueryByID(context.Background(), "123")
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedUser, usr)

    // Verify the mocks were called
    mockCacheStorer.AssertExpectations(t)
}

func TestCore_QueryByIDNoCache(t *testing.T) {
    // Setup
    mockCacheStorer := new(MockCacheStorer)
    mockStorer := new(MockStorer)
    core, _ := user.NewCore(mockStorer, nil, mockCacheStorer)

    expectedUser := user.User{
        ID:    primitive.NewObjectID(),
        Name:  "Jane Doe",
        Email: "jane.doe@example.com",
    }

    // Simulate a cache miss by returning redis.Nil
    mockCacheStorer.On("Get", mock.Anything, "user:123", mock.AnythingOfType("*user.User")).
        Return(redis.Nil).
        Once()

    // Simulate the database call after cache miss
    mockStorer.On("QueryByID", mock.Anything, "123").
        Return(expectedUser, nil).
        Once()

    // Test
    usr, err := core.QueryByID(context.Background(), "123")

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedUser, usr)

    // Verify the mocks were called
    mockCacheStorer.AssertExpectations(t)
    mockStorer.AssertExpectations(t)
}

// TestCore_Create verifies that the Create function correctly creates a user, stores it in the DB, and caches it.
func TestCore_Create(t *testing.T) {
    // Setup
    mockCacheStorer := new(MockCacheStorer)
    mockStorer := new(MockStorer)
    mockEthClient := new(blockchain.MockClient)
    core, _ := user.NewCore(mockStorer, mockEthClient, mockCacheStorer)

    // Create a new user
    email, _ := mail.ParseAddress("new.user@example.com")
    newUser := user.User{
        Name:  "New User",
        Email: email.String(),
    }

    expectedUser := newUser
    expectedUser.ID = primitive.NewObjectID()

    // Mock the Eth client to return account details
    mockEthClient.On("CreateEthAccount").
        Return(blockchain.AccountData{
            PublicKey:   []byte("publicKey"),
            PrivateKey:  []byte("privateKey"),
            AddressHex:  "0xdf56089fedbacf7ba0bc217dafbffa2fc08a93fdc68e1e42411a14efcf23636e",
        }, nil).
        Once()

    // Mock the store to insert the user
    mockStorer.On("Create", mock.Anything, mock.AnythingOfType("user.User")).
        Return(expectedUser, nil).
        Once()

    // Mock cache set
    mockCacheStorer.On("Set", mock.Anything, expectedUser).
        Return(nil).
        Once()

    // Test
    createdUser, err := core.Create(context.Background(), newUser)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedUser.Email, createdUser.Email)

    // Verify the mocks were called
    mockStorer.AssertExpectations(t)
    mockCacheStorer.AssertExpectations(t)
    mockEthClient.AssertExpectations(t)
}

// TestCore_Delete verifies that the Delete function correctly removes a user from the DB and cache.
func TestCore_Delete(t *testing.T) {
    // Setup
    mockCacheStorer := new(MockCacheStorer)
    mockStorer := new(MockStorer)
    core, _ := user.NewCore(mockStorer, nil, mockCacheStorer)

    usr := user.User{
        ID:    primitive.NewObjectID(),
        Name:  "User To Delete",
        Email: "delete.me@example.com",
    }

    // Mock the store delete
    mockStorer.On("Delete", mock.Anything, usr).
        Return(nil).
        Once()

    // Mock cache delete
    mockCacheStorer.On("Delete", mock.Anything, usr.CacheKey()).
        Return(nil).
        Once()

    // Test
    err := core.Delete(context.Background(), usr)

    // Assert
    assert.NoError(t, err)

    // Verify the mocks were called
    mockStorer.AssertExpectations(t)
    mockCacheStorer.AssertExpectations(t)
}

// TestCore_Count verifies that the Count function correctly returns the total number of users.
func TestCore_Count(t *testing.T) {
    // Setup
    mockStorer := new(MockStorer)
    core, _ := user.NewCore(mockStorer, nil, nil)

    filter := user.QueryFilter{}
    expectedCount := 5

    // Mock the store count
    mockStorer.On("Count", mock.Anything, filter).
        Return(expectedCount, nil).
        Once()

    // Test
    count, err := core.Count(context.Background(), filter)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedCount, count)

    // Verify the mocks were called
    mockStorer.AssertExpectations(t)
}

// TestCore_QueryByEmail verifies that the QueryByEmail function retrieves the correct user.
func TestCore_QueryByEmail(t *testing.T) {
    // Setup
    mockStorer := new(MockStorer)
    core, _ := user.NewCore(mockStorer, nil, nil)

    email, _ := mail.ParseAddress("query.by.email@example.com")
    expectedUser := user.User{
        ID:    primitive.NewObjectID(),
        Name:  "Query User",
        Email: email.String(),
    }

    // Mock the store query by email
    mockStorer.On("QueryByEmail", mock.Anything, *email).
        Return(expectedUser, nil).
        Once()

    // Test
    usr, err := core.QueryByEmail(context.Background(), *email)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedUser.Email, usr.Email)

    // Verify the mocks were called
    mockStorer.AssertExpectations(t)
}

// TestCore_Authenticate verifies that the Authenticate function checks the email and password correctly.
func TestCore_Authenticate(t *testing.T) {
    // Setup
    mockStorer := new(MockStorer)
    core, _ := user.NewCore(mockStorer, nil, nil)

    email, _ := mail.ParseAddress("auth.user@example.com")
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
    expectedUser := user.User{
        ID:           primitive.NewObjectID(),
        Name:         "Auth User",
        Email:        email.String(),
        PasswordHash: hashedPassword,
    }

    // Mock the store query by email
    mockStorer.On("QueryByEmail", mock.Anything, *email).
        Return(expectedUser, nil).
        Once()

    // Test
    usr, err := core.Authenticate(context.Background(), *email, "password")

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedUser.Email, usr.Email)

    // Verify the mocks were called
    mockStorer.AssertExpectations(t)
}


