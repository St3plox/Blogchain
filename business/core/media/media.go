// Package media handles media for the posts
package media

import (
	"context"
	"errors"
	"fmt"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	"github.com/St3plox/Blogchain/foundation/cachestore"
	"github.com/redis/go-redis/v9"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound     = errors.New("media not found")
	ErrUserNotOwner = errors.New("user does not own the media")
)

type Storer interface {
	Create(ctx context.Context, media Media) (Media, error)
	CreateMultiple(ctx context.Context, media []Media) ([]Media, error)
	Update(ctx context.Context, media Media) (Media, error)
	Delete(ctx context.Context, media Media) error
	DeleteByID(ctx context.Context, mediaID string) error
	QueryByID(ctx context.Context, mediaID string) (Media, error)
	QueryByIDs(ctx context.Context, mediaIDs []string) ([]Media, error)
}

// Core manages the set of APIs for media access.
type Core struct {
	storer      Storer
	cacheStorer cachestore.CacheStorer
	userCore    *user.Core

	MaxFileSizeMb int64
}

func NewCore(storer Storer, cacheStorer cachestore.CacheStorer, userCore *user.Core) *Core {
	return &Core{
		storer:      storer,
		cacheStorer: cacheStorer,
		userCore:    userCore,
	}
}

func (c *Core) Create(ctx context.Context, newMedia NewMedia) (MediaData, error) {
	claims := auth.GetClaims(ctx)

	user, err := c.userCore.QueryByID(ctx, claims.Subject)
	if err != nil {
		return MediaData{}, fmt.Errorf("media core create get user: %w", err)
	}

	media := Media{
		Filename:  newMedia.Filename,
		Length:    newMedia.Length,
		FileBytes: newMedia.FileBytes,
		OwnerID:   user.ID,
	}

	media, err = c.storer.Create(ctx, media)
	if err != nil {
		return MediaData{}, fmt.Errorf("core error create: %w", err)
	}

	return MapTo(media), nil
}

func (c *Core) CreateMultiple(ctx context.Context, newMediaList []NewMedia) ([]MediaData, error) {
	claims := auth.GetClaims(ctx)

	user, err := c.userCore.QueryByID(ctx, claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("media core create get user: %w", err)
	}

	var mediaList []Media
	for _, newMedia := range newMediaList {
		media := Media{
			Filename:  newMedia.Filename,
			Length:    newMedia.Length,
			FileBytes: newMedia.FileBytes,
			OwnerID:   user.ID,
		}
		mediaList = append(mediaList, media)
	}

	mediaList, err = c.storer.CreateMultiple(ctx, mediaList)
	if err != nil {
		return nil, fmt.Errorf("core error create: %w", err)
	}

	return MapToMultiple(mediaList), nil
}

func (c *Core) Delete(ctx context.Context, media Media) error {
	err := c.storer.Delete(ctx, media)
	if err != nil {
		if err == ErrNotFound {
			return err
		}
		return fmt.Errorf("core error delete: %w", err)
	}

	err = c.cacheStorer.Delete(ctx, IdToCacheKey(media.ID.Hex()))
	if err != nil {
		return fmt.Errorf("core error cache delete: %w", err)
	}

	return nil
}

func (c *Core) DeleteByID(ctx context.Context, mediaID string) error {
	claims := auth.GetClaims(ctx)

	user, err := c.userCore.QueryByID(ctx, claims.Subject)
	if err != nil {
		return fmt.Errorf("media core error get user: %w", err)
	}

	media, err := c.storer.QueryByID(ctx, mediaID)
	if err != nil {
		if err == ErrNotFound {
			return err
		}
		return fmt.Errorf("core error delete, query error: %w", err)
	}

	if media.OwnerID != user.ID {
		return ErrUserNotOwner
	}

	err = c.storer.DeleteByID(ctx, mediaID)
	if err != nil {
		return fmt.Errorf("core delete error %w", err)
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
		return Media{}, fmt.Errorf("core error get: %w", err)
	}

	c.cacheStorer.Set(ctx, media)

	return media, nil
}

// QueryByIDs retrieves multiple media records by their IDs, checking the cache first and querying the database for any missing records.
func (c *Core) QueryByIDs(ctx context.Context, mediaIDs []string) ([]Media, error) {
	var medias []Media
	var idsToQuery []string

	for _, id := range mediaIDs {
		var media Media
		err := c.cacheStorer.Get(ctx, IdToCacheKey(id), &media)
		if err == nil {
			medias = append(medias, media)
		} else if err == redis.Nil {
			idsToQuery = append(idsToQuery, id)
		} else {
			return nil, fmt.Errorf("core cache get by ids: %w", err)
		}
	}

	if len(idsToQuery) > 0 {
		dbMedias, err := c.storer.QueryByIDs(ctx, idsToQuery)
		if err != nil {
			return nil, fmt.Errorf("core error get by ids: %w", err)
		}

		for _, media := range dbMedias {
			medias = append(medias, media)
			if cacheErr := c.cacheStorer.Set(ctx, media); cacheErr != nil {
				return nil, fmt.Errorf("core cache set: %w", cacheErr)
			}
		}
	}

	return medias, nil
}
