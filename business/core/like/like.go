package like

import (
	"context"
	"fmt"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/foundation/cachestore"
)

type Storer interface {
	Create(ctx context.Context, newLike Like) (Like, error)
	QueryByID(ctx context.Context, likeID string) (Like, error)
	QueryAllByUserID(ctx context.Context, userID string) ([]Like, error)
	QueryAllByPostID(ctx context.Context, userID string) ([]Like, error)
	Update(ctx context.Context, updatedLike Like) (Like, error)
	DeleteByID(ctx context.Context, likeID string) error
}

type Core struct {
	cache    cachestore.CacheStore
	storer   Storer
	userSore user.Storer
}

func NewCore(cache cachestore.CacheStore, storer Storer, userStore user.Storer) *Core {
	return &Core{
		cache:    cache,
		storer:   storer,
		userSore: userStore,
	}
}

func (c *Core) Create(ctx context.Context, newLike Like) (Like, error) {
	claims := auth.GetClaims(ctx)

	user, err := c.userSore.QueryByID(ctx, claims.Subject)
	if err != nil {
		return Like{}, err
	}

	newLike.UserID = user.ID

	savedLike, err := c.storer.Create(ctx, newLike)
	if err != nil {
		return Like{}, fmt.Errorf("error creating like: %w", err)
	}

	// Cache the new like
	if err := c.cache.Set(ctx, savedLike); err != nil {
		return Like{}, fmt.Errorf("error caching like: %w", err)
	}

	return savedLike, nil
}

func (c *Core) QueryByID(ctx context.Context, likeID string) (Like, error) {
	var like Like

	cacheKey := IdToCacheKey(likeID)
	if err := c.cache.Get(ctx, cacheKey, &like); err == nil {
		return like, nil
	}

	like, err := c.storer.QueryByID(ctx, likeID)
	if err != nil {
		return Like{}, fmt.Errorf("error querying like by ID: %w", err)
	}

	if err := c.cache.Set(ctx, like); err != nil {
		return Like{}, fmt.Errorf("error caching like: %w", err)
	}

	return like, nil
}

// QueryAllByPostID retrieves all likes for a specific post
func (c *Core) QueryAllByPostID(ctx context.Context, postID string) ([]Like, error) {
	// Fetch likes from storer
	likes, err := c.storer.QueryAllByPostID(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("error querying all likes by post ID: %w", err)
	}

	return likes, nil
}

func (c *Core) QueryAllByUserID(ctx context.Context, userID string) ([]Like, error) {

	likes, err := c.storer.QueryAllByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying all likes by user ID: %w", err)
	}

	return likes, nil
}

func (c *Core) Update(ctx context.Context, updatedLike Like) (Like, error) {

	//TODO: add check that this user updates like
	updatedLike, err := c.storer.Update(ctx, updatedLike)
	if err != nil {
		return Like{}, fmt.Errorf("error updating like: %w", err)
	}

	// Update cache
	if err := c.cache.Set(ctx, updatedLike); err != nil {
		return Like{}, fmt.Errorf("error updating like cache: %w", err)
	}

	return updatedLike, nil
}

func (c *Core) DeleteByID(ctx context.Context, likeID string) error {
	cacheKey := IdToCacheKey(likeID)

	if err := c.cache.Delete(ctx, cacheKey); err != nil {
		return fmt.Errorf("error deleting like from cache: %w", err)
	}

	if err := c.storer.DeleteByID(ctx, likeID); err != nil {
		return fmt.Errorf("error deleting like by ID: %w", err)
	}

	return nil
}