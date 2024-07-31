// Package media handles media for the posts
package media

import (
	"context"
	"errors"
	"fmt"

	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/redis/go-redis/v9"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound = errors.New("media not found")
)

type Storer interface {
	Create(ctx context.Context, media Media) (Media, error)
	Update(ctx context.Context, media Media) (Media, error)
	Delete(ctx context.Context, media Media) error
	DeleteByID(ctx context.Context, mediaID string) error
	QueryByID(ctx context.Context, mediaID string) (Media, error)
}

// Core manages the set of APIs for media access.
type Core struct {
	storer      Storer
	cacheStorer cachestore.CacheStorer

	MaxFileSizeMb int64
}

func NewCore(storer Storer, cacheStorer cachestore.CacheStorer) *Core {
	return &Core{
		storer:      storer,
		cacheStorer: cacheStorer,
	}
}

func (c *Core) Create(ctx context.Context, newMedia NewMedia) (Media, error) {

	media := Media{
		Filename:  newMedia.Filename,
		Length:    newMedia.Length,
		FileBytes: newMedia.FileBytes,
	}

	media, err := c.storer.Create(ctx, media)
	if err != nil {
		return Media{}, fmt.Errorf("core error create: %w", err)
	}

	return media, nil
}

func (c *Core) Update(ctx context.Context, media Media) (Media, error) {

	media, err := c.storer.Update(ctx, media)
	if err != ErrNotFound {
		return Media{}, fmt.Errorf("core error update: %w", err)
	} else if err != nil {
		return Media{}, err
	}

	err = c.cacheStorer.Set(ctx, media)
	if err != nil {
		return Media{}, fmt.Errorf("core error update cache set: %w", err)
	}

	return media, nil
}

func (c *Core) Delete(ctx context.Context, media Media) error {

	err := c.storer.Delete(ctx, media)
	if err != ErrNotFound {
		return fmt.Errorf("core error delete: %w", err)
	} else if err != nil {
		return err
	}

	err = c.cacheStorer.Delete(ctx, IdToCacheKey(media.ID.Hex()))
	if err != nil {
		return fmt.Errorf("core error cahe delete: %w", err)
	}

	return nil
}

func (c *Core) DeleteByID(ctx context.Context, mediaID string) error {

	err := c.storer.DeleteByID(ctx, mediaID)
	if err != ErrNotFound {
		return fmt.Errorf("core error delete: %w", err)
	} else if err != nil {
		return err
	}

	err = c.cacheStorer.Delete(ctx, IdToCacheKey(mediaID))
	if err != nil {
		return fmt.Errorf("core error cache delete: %w", err)
	}

	return nil
}

func (c *Core) QueryByID(ctx context.Context, mediaID string) (Media, error) {

	var media Media
	err := c.cacheStorer.Get(ctx, IdToCacheKey(mediaID), &media)
	if err != nil {
		if err != redis.Nil {
			return Media{}, fmt.Errorf("core cache get: %w", err)
		}
	}

	media, err = c.storer.QueryByID(ctx, mediaID)
	if err != nil {
		if err != ErrNotFound {
			return Media{}, err
		}
		return  Media{}, fmt.Errorf("core error get: %w", err) 
	}
	
	return media, nil
}
