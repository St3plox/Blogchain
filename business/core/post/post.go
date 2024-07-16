package post

import (
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/St3plox/Blogchain/foundation/blockchain/auth"
	"github.com/St3plox/Blogchain/foundation/blockchain/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Core struct {
	postContract *contract.PostContract
	admin        *auth.Admin
}

func NewCore(postContract *contract.PostContract, admin *auth.Admin) *Core {
	return &Core{
		postContract: postContract,
		admin:        admin,
	}
}

type Storer interface {
	Create(ctx context.Context, post Post, userPrivateKey []byte) error
	QueryByAddress(ctx context.Context, userAddress string) ([]Post, error)
	GetPostByIndex(ctx context.Context, userAddress string, index uint64) (Post, error)
	GetAllPostsSorted(ctx context.Context) ([]Post, error)
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
	tx, err := c.postContract.Contract.Post(auth, np.Title, np.Content, uint8(np.Category), address)
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

	fmt.Println("Creating post with:")
	fmt.Printf("Title: %s, Content: %s, Author: %s\n", np.Title, np.Content, userAddressHex)

	var newPost Post
	for _, log := range receipt.Logs {
		event, err := c.postContract.Contract.ParsePostPublished(*log)
		if err == nil {
			newPost = Post{
				Author:    event.Author,
				Title:     event.Title,
				Category:  Category(event.Category),
				Timestamp: big.NewInt(receipt.BlockNumber.Int64()),
				Content:   np.Content,
			}
			break
		} else {
			fmt.Println("Error parsing event:", err)
		}
	}

	return newPost, nil
}

func (c *Core) QueryByAddress(ctx context.Context, userAddressHex string) ([]Post, error) {

	if !common.IsHexAddress(userAddressHex) {
		return nil, fmt.Errorf("invalid address: %s", userAddressHex)
	}

	address := common.HexToAddress(userAddressHex)

	posts, err := c.postContract.Contract.GetUsersPost(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}

	var result []Post
	for _, post := range posts {
		result = append(result, Post{
			Author:    post.Author,
			Title:     post.Title,
			Content:   post.Content,
			Timestamp: post.Timestamp,
			Category:  Category(post.Category),
		})
	}

	return result, nil
}

func (c *Core) GetPostByIndex(ctx context.Context, userAddressHex string, index uint64) (Post, error) {

	address := common.HexToAddress(userAddressHex)

	post, err := c.postContract.Contract.GetPostByIndex(&bind.CallOpts{Context: ctx}, address, new(big.Int).SetUint64(index))
	if err != nil {
		return Post{}, err
	}

	return Post{
		Author:    post.Author,
		Title:     post.Title,
		Content:   post.Content,
		Timestamp: post.Timestamp,
		Category:  Category(post.Category),
	}, nil
}

func (c *Core) GetAllPostsSorted(ctx context.Context) ([]Post, error) {
	posts, err := c.postContract.Contract.GetAllPosts(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	var result []Post
	for _, post := range posts {
		result = append(result, Post{
			Author:    post.Author,
			Title:     post.Title,
			Content:   post.Content,
			Timestamp: post.Timestamp,
			Category:  Category(post.Category),
		})
	}

	// Sort all posts by timestamp in descending order
	sort.Slice(result, func(i, j int) bool {
		return result[i].Timestamp.Cmp(result[j].Timestamp) > 0
	})

	return result, nil
}
