package post_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/St3plox/Blogchain/business/core/post"
	"github.com/St3plox/Blogchain/foundation/blockchain/auth"
	"github.com/St3plox/Blogchain/foundation/blockchain/contract"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// WARNING: These accounts, and their private keys, are publicly known.
// Any funds sent to them on Mainnet or any other live network WILL BE LOST.
// NOTE: took it from hardhat
const adminKey = "0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"
const testnetUrl = "http://127.0.0.1:8545/"
const userAdressHex = "0xFABB0ac9d68B0B445fB7357272Ff202C5651694a"

func initCore(mockCache cachestore.CacheStore, ctx context.Context) (*post.Core, error) {

	// Initialize Ethereum client
	client, err := ethclient.Dial(testnetUrl)
	if err != nil {
		return nil, err
	}

	admin, err := auth.NewAdmin(adminKey, client)
	if err != nil {
		return nil, err
	}
	cAuth, err := admin.GenerateAuth(ctx)
	if err != nil {
		return nil, err
	}

	cAuth.GasLimit = 6000000

	_, _, instance, err := contract.DeployContract(cAuth, client)
	if err != nil {
		return nil, err
	}

	postContract, err := contract.NewPostContract(client, instance)
	if err != nil {
		return nil, err
	}

	return post.NewCore(postContract, admin, mockCache), nil
}

func TestCore_QueryByIDFromCahe(t *testing.T) {

	ctx := context.Background()
	mockCache := new(cachestore.MockCacheStore)

	core, err := initCore(mockCache, ctx)
	if err != nil {
		t.Fatalf("Failed to create core: %v", err)
	}

	postId := big.NewInt(1)
	cachedPost := post.Post{
		ID:        postId,
		Title:     "Cached Post",
		Content:   "This is a cached post",
		Category:  post.Category(1),
		Timestamp: big.NewInt(1234567890),
		Tags:      []string{"cached"},
	}

	// Mock cache get
	mockCache.On("Get", ctx, post.IdToCacheKey(postId.String()), mock.AnythingOfType("*post.Post")).Run(func(args mock.Arguments) {
		val := args.Get(2).(*post.Post)
		*val = cachedPost
	}).Return(nil)

	queriedPost, err := core.QueryById(ctx, postId)
	if err != nil {
		t.Fatalf("Failed to query post by ID: %v", err)
	}

	assert.Equal(t, cachedPost.ID, queriedPost.ID, "Post ID should match")
	assert.Equal(t, cachedPost.Title, queriedPost.Title, "Post title should match")
	assert.Equal(t, cachedPost.Content, queriedPost.Content, "Post content should match")

	// Check cache interaction
	mockCache.AssertExpectations(t)
}

// TestCore_QueryByIDNoCache tests querying a post by ID where no cache is involved
// TestCore_QueryByIDNoCache tests querying a post by ID where no cache is involved
func TestCore_QueryByIDNoCache(t *testing.T) {
	ctx := context.Background()
	mockCache := new(cachestore.MockCacheStore)

	core, err := initCore(mockCache, ctx)
	if err != nil {
		t.Fatalf("Failed to create core: %v", err)
	}

	// Create a post to query
	newPost := post.NewPost{
		Title:      "Test Post",
		Content:    "This is a test post",
		Category:   post.Category(1),
		Tags:       []string{"test", "post"},
		MediaNames: []string{},
		MediaUrls:  []string{},
	}

	// Mock cache get to simulate no cache
	mockCache.On("Get", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*post.Post")).Return(redis.Nil)

	// Mock cache set to simulate successful cache storage
	mockCache.On("Set", ctx, mock.AnythingOfType("*post.Post")).Return(nil)

	// Create the post
	createdPost, err := core.Create(ctx, newPost, userAdressHex)
	if err != nil {
		t.Fatalf("Failed to create post: %v", err)
	}

	// Mock cache get to simulate no cache again during query
	mockCache.On("Get", ctx, post.IdToCacheKey(createdPost.ID.String()), mock.AnythingOfType("*post.Post")).Return(redis.Nil)

	// Query by ID
	queriedPost, err := core.QueryById(ctx, createdPost.ID)
	if err != nil {
		t.Fatalf("Failed to query post by ID: %v", err)
	}

	// Verify that the queried post matches the created post
	assert.Equal(t, createdPost.ID, queriedPost.ID, "Post ID should match")
	assert.Equal(t, createdPost.Title, queriedPost.Title, "Post title should match")
	assert.Equal(t, createdPost.Content, queriedPost.Content, "Post content should match")

	// Check cache interactions
	mockCache.AssertExpectations(t)
}
