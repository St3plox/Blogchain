package like

import (
	"context"

	"github.com/St3plox/Blogchain/foundation/cachestore"
)

type Storer interface {
	Create(ctx context.Context, newLike Like) (Like, error)
	QueryByID(ctx context.Context, likeID string) (Like, error)
	QueryAllByUserID(ctx context.Context, likeID string) ([]Like, error)
	Update(ctx context.Context, updatedLike Like) (Like, error)
	DeleteByID(ctx context.Context, likeID string) error
}

type Core struct {
	cache  cachestore.CacheStore
	storer Storer
}

func NewCore(cache cachestore.CacheStore, storer Storer) *Core {
	return &Core{
		cache:  cache,
		storer: storer,
	}
}
