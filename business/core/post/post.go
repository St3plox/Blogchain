package post

import (
	"context"
	"fmt"
	"math/big"

	"github.com/St3plox/Blogchain/foundation/blockchain/auth"
	"github.com/St3plox/Blogchain/foundation/blockchain/contract"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
)


type Core struct {
	postContract *contract.PostContract
	admin        auth.ContractSigner
	cacheStorer  cachestore.CacheStore
}

func NewCore(postContract *contract.PostContract, admin auth.ContractSigner, cacheStorer cachestore.CacheStore) *Core {
	return &Core{
		postContract: postContract,
		admin:        admin,
		cacheStorer:  cacheStorer,
	}
}

type Storer interface {
	Create(ctx context.Context, np NewPost, userAddressHex string) (Post, error)
	Query(ctx context.Context, page uint64, pageSize uint64) ([]Post, error)
	QueryByAddress(ctx context.Context, userAddressHex string, page uint64, pageSize uint64) ([]Post, error)
	QueryByIndex(ctx context.Context, userAddressHex string, index uint64) (Post, error)
	QueryById(ctx context.Context, id *big.Int) (Post, error)
}

func (c *Core) Create(ctx context.Context, np NewPost, userAddressHex string) (Post, error) {

	if !common.IsHexAddress(userAddressHex) {
		return Post{}, fmt.Errorf("invalid address: %s", userAddressHex)
	}

	address := common.HexToAddress(userAddressHex)
	fmt.Println(userAddressHex)

	auth, err := c.admin.GenerateAuth(ctx)
	if err != nil {
		return Post{}, err
	}
	tx, err := c.postContract.Contract.Post(auth, np.Title, np.Content, uint8(np.Category), np.Tags, address, np.MediaNames, np.MediaUrls)
	if err != nil {
		return Post{}, fmt.Errorf("error making post: %e", err)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, c.postContract.Client, tx)
	if err != nil {
		return Post{}, err
	}
	if receipt.Status != 1 {
		return Post{}, fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	var newPost Post
	for _, log := range receipt.Logs {
		event, err := c.postContract.Contract.ParsePostPublished(*log)
		if err == nil {
			newPost = Post{
				ID:        event.Id,
				Author:    event.Author,
				Title:     event.Title,
				Category:  Category(event.Category),
				Timestamp: big.NewInt(receipt.BlockNumber.Int64()),
				Content:   np.Content,
				Tags:      event.Tags,
			}
			break
		} else {
			fmt.Println("Error parsing event:", err)
		}
	}

	err = c.cacheStorer.Set(ctx, &newPost)
	if err != nil {
		return Post{}, fmt.Errorf("cache error: %w", err)
	}

	return newPost, nil
}

func (c *Core) QueryByAddress(ctx context.Context, userAddressHex string, page uint64, pageSize uint64) ([]Post, error) {

	if !common.IsHexAddress(userAddressHex) {
		return nil, fmt.Errorf("invalid address: %s", userAddressHex)
	}

	address := common.HexToAddress(userAddressHex)

	posts, err := c.postContract.Contract.GetUsersPostPaginated(
		&bind.CallOpts{Context: ctx},
		address,
		new(big.Int).SetUint64(page),
		new(big.Int).SetUint64(pageSize),
	)

	if err != nil {
		return nil, err
	}

	var result []Post
	for _, post := range posts {
		result = append(result, mapTo(post))
	}

	return result, nil
}

func (c *Core) QueryByIndex(ctx context.Context, userAddressHex string, index uint64) (Post, error) {
	address := common.HexToAddress(userAddressHex)

	post, err := c.postContract.Contract.GetPostByIndex(&bind.CallOpts{Context: ctx}, address, new(big.Int).SetUint64(index))
	if err != nil {
		return Post{}, err
	}

	return mapTo(post), nil
}

func (c *Core) QueryById(ctx context.Context, id *big.Int) (Post, error) {

	{ // Checking for cached post
		var post Post
		err := c.cacheStorer.Get(ctx, IdToCacheKey(id.String()), &post)
		if err == nil {
			return post, nil
		} else if err != redis.Nil {
			return Post{}, fmt.Errorf("cache get: %w", err)
		}
	}

	post, err := c.postContract.Contract.GetPostByID(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return Post{}, err
	}

	postRes := mapTo(post)

	c.cacheStorer.Set(ctx, &postRes)

	return postRes, nil
}

func (c *Core) Query(ctx context.Context, page uint64, pageSize uint64) ([]Post, error) {
	// Retrieve all posts to calculate the total number of posts
	posts, err := c.postContract.Contract.GetAllPosts(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	totalPostsCount := uint64(len(posts))

	// Calculate the start and end indices for the paginated results
	start := page * pageSize
	if start >= totalPostsCount {
		return nil, fmt.Errorf("page out of range")
	}

	end := start + pageSize
	if end > totalPostsCount {
		end = totalPostsCount
	}

	// Slice the posts array to get only the posts for the current page
	paginatedPosts := posts[start:end]

	// Convert the paginated posts to the Post type used by Go
	var result []Post
	for _, post := range paginatedPosts {
		postRes := mapTo(post)
		c.cacheStorer.Set(ctx, &postRes)
		result = append(result, mapTo(post))
	}

	return result, nil
}
