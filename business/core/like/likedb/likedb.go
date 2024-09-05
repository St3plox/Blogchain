package likedb

import (
	"context"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName         = "blogchainDB"
	collectionName = "likeCollection"
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

// Create inserts a new Like document into the collection.
func (s *Store) Create(ctx context.Context, newLike like.Like) (like.Like, error) {

	newLike.ID = primitive.NewObjectID()

	_, err := s.collection.InsertOne(ctx, newLike)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to create Like")
		return like.Like{}, ErrLikeCreationFailed
	}
	return newLike, nil
}

// QueryByID retrieves a Like document by its ID.
func (s *Store) QueryByID(ctx context.Context, likeID string) (like.Like, error) {
	objectID, err := primitive.ObjectIDFromHex(likeID)
	if err != nil {
		s.log.Error().Err(err).Msg("Invalid Like ID format")
		return like.Like{}, ErrInvalidIDFormat
	}

	var foundLike like.Like
	err = s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&foundLike)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return like.Like{}, ErrLikeNotFound
		}
		s.log.Error().Err(err).Msg("Failed to query Like by ID")
		return like.Like{}, ErrLikeQueryFailed
	}
	return foundLike, nil
}

// QueryAllByUserID retrieves all Like documents for a specific user.
func (s *Store) QueryAllByUserID(ctx context.Context, userID string) ([]like.Like, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		s.log.Error().Err(err).Msg("Invalid User ID format")
		return nil, ErrInvalidUserIDFormat
	}

	cursor, err := s.collection.Find(ctx, bson.M{"user_id": objectID})
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to query Likes by User ID")
		return nil, err
	}
	defer cursor.Close(ctx)

	var likes []like.Like
	if err = cursor.All(ctx, &likes); err != nil {
		s.log.Error().Err(err).Msg("Failed to decode Likes")
		return nil, err
	}
	return likes, nil
}

// Update modifies an existing Like document.
func (s *Store) Update(ctx context.Context, updatedLike like.Like) (like.Like, error) {
	filter := bson.M{"_id": updatedLike.ID}
	update := bson.M{"$set": updatedLike}

	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to update Like")
		return like.Like{}, ErrLikeUpdateFailed
	}
	if result.MatchedCount == 0 {
		return like.Like{}, ErrLikeNotFound
	}
	return updatedLike, nil
}

// DeleteByID removes a Like document by its ID.
func (s *Store) DeleteByID(ctx context.Context, likeID string) error {
	objectID, err := primitive.ObjectIDFromHex(likeID)
	if err != nil {
		s.log.Error().Err(err).Msg("Invalid Like ID format")
		return ErrInvalidIDFormat
	}

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to delete Like by ID")
		return ErrLikeDeletionFailed
	}
	if result.DeletedCount == 0 {
		return ErrLikeNotFound
	}
	return nil
}

// QueryAllByPostID retrieves all Like documents for a specific post.
func (s *Store) QueryAllByPostID(ctx context.Context, postID string) ([]like.Like, error) {

	cursor, err := s.collection.Find(ctx, bson.M{"post_id": postID})
	if err != nil {
		s.log.Error().Err(err).Msg("Failed to query Likes by Post ID")
		return nil, err
	}
	defer cursor.Close(ctx)

	var likes []like.Like
	if err = cursor.All(ctx, &likes); err != nil {
		s.log.Error().Err(err).Msg("Failed to decode Likes")
		return nil, err
	}
	return likes, nil
}
