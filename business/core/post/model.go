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
	ID        *big.Int
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  Category
}

type NewPost struct {
	Title    string
	Content  string
	Category Category
}

func (p Post) IsEmpty() bool {
	return reflect.DeepEqual(p, Post{})
}
