package post

import (
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

type Category uint8

const (
	Blog Category = iota
	News
	Article
)


type Post struct {
	ID        *big.Int        `json:"id"`
	Author    common.Address  `json:"author"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Timestamp *big.Int        `json:"timestamp"`
	Category  Category        `json:"category"`
	Tags      []string        `json:"tags"`
}

type NewPost struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category Category `json:"category"`
	Tags     []string `json:"tags"`
}

func (p Post) IsEmpty() bool {
	return reflect.DeepEqual(p, Post{})
}
