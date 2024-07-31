package mediadb

import (
	"context"
	"fmt"

	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName         = "blogchainDB"
	collectionName = "mediaCollection"
)

type Store struct {
	log        *zerolog.Logger
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStore(log *zerolog.Logger, client *mongo.Client) *Store {
	return &Store{
		log:        log,
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (s *Store) Create(ctx context.Context, m media.Media) (media.Media, error) {
	result, err := s.collection.InsertOne(ctx, m)
	if err != nil {
		return media.Media{}, fmt.Errorf("create: %w", err)
	}

	m.ID = result.InsertedID.(primitive.ObjectID)
	return m, nil
}

func (s *Store) Update(ctx context.Context, m media.Media) (media.Media, error) {
	filter := bson.M{"_id": m.ID}
	update := bson.M{"$set": m}

	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return media.Media{}, fmt.Errorf("update: %w", err)
	}

	if result.MatchedCount == 0 {
		return media.Media{}, media.ErrNotFound
	}

	return m, nil
}

func (s *Store) Delete(ctx context.Context, m media.Media) error {
	filter := bson.M{"_id": m.ID}

	result, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	if result.DeletedCount == 0 {
		return media.ErrNotFound
	}

	return nil
}

func (s *Store) DeleteByID(ctx context.Context, mediaID string) error {
	id, err := primitive.ObjectIDFromHex(mediaID)
	if err != nil {
		return fmt.Errorf("delete by id: %w", err)
	}

	filter := bson.M{"_id": id}

	result, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete by id: %w", err)
	}

	if result.DeletedCount == 0 {
		return media.ErrNotFound
	}

	return nil
}

func (s *Store) QueryByID(ctx context.Context, mediaID string) (media.Media, error) {
	id, err := primitive.ObjectIDFromHex(mediaID)
	if err != nil {
		return media.Media{}, fmt.Errorf("query by id: %w", err)
	}

	filter := bson.M{"_id": id}

	var m media.Media
	err = s.collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return media.Media{}, media.ErrNotFound
		}
		return media.Media{}, fmt.Errorf("query by id: %w", err)
	}

	return m, nil
}

func (s *Store) QueryByIDs(ctx context.Context, mediaIDs []string) ([]media.Media, error) {
	var ids []primitive.ObjectID
	for _, id := range mediaIDs {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("invalid id %s: %w", id, err)
		}
		ids = append(ids, objID)
	}

	filter := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get by ids: %w", err)
	}
	defer cursor.Close(ctx)

	var medias []media.Media
	for cursor.Next(ctx) {
		var m media.Media
		if err := cursor.Decode(&m); err != nil {
			return nil, fmt.Errorf("decode media: %w", err)
		}
		medias = append(medias, m)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return medias, nil
}
