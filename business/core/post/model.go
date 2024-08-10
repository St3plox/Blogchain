package post

import (
	"math/big"
	"reflect"
	"time"

	"github.com/St3plox/Blogchain/foundation/blockchain/contract"
	"github.com/ethereum/go-ethereum/common"
)

type Category uint8

const (
	Blog Category = iota
	News
	Article
)

type Post struct {
	ID        *big.Int          `json:"id"`
	Author    common.Address    `json:"author"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	Timestamp *big.Int          `json:"timestamp"`
	Category  Category          `json:"category"`
	Tags      []string          `json:"tags"`
	Media     map[string]string `json:"media"` //[name]url
}

type NewPost struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Category   Category `json:"category"`
	Tags       []string `json:"tags"`
	MediaNames []string `json:"media_names"`
	MediaUrls  []string `json:"media_urls"`
}

func (p *Post) CacheKey() string {
	return IdToCacheKey(p.ID.String())
}

func (p *Post) CacheExpiration() time.Duration {
	return 24 * time.Hour
}

func IdToCacheKey(id string) string {
	return "post:" + id
}

func (p Post) IsEmpty() bool {
	return reflect.DeepEqual(p, Post{})
}

func mapTo(postStoragePost contract.PostStoragePost) Post {

	mediaUrls := make(map[string]string)
	for i, name := range postStoragePost.MediaNames {
		mediaUrls[name] = postStoragePost.MediaUrls[i]
	}

	return Post{
		ID:        postStoragePost.Id,
		Author:    postStoragePost.Author,
		Title:     postStoragePost.Title,
		Content:   postStoragePost.Content,
		Timestamp: postStoragePost.Timestamp,
		Tags:      postStoragePost.Tags,
		Category:  Category(postStoragePost.Category),
		Media:     mediaUrls,
	}
}
