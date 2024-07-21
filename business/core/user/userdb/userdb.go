package userdb

import (
	"context"
	"net/mail"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/data/order"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "userDB"
const collectionName = "userCollection"

type Store struct {
	log        *zerolog.Logger
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStore(log *zerolog.Logger, client *mongo.Client) *Store {
	store := &Store{
		log:        log,
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}

	// Create unique indexes for email and name
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := store.collection.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create indexes")
	}

	return store
}

// NOTE: add email verification in the future
// NOTE: add email verification in the future
func (s *Store) Create(ctx context.Context, usr user.User) (user.User, error) {
	res, err := s.collection.InsertOne(ctx, usr)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			s.log.Error().Err(err).Msg("mongodb: duplicate key error")
			return user.User{}, err // Handle duplicate key error appropriately
		}
		s.log.Error().Err(err).Msg("mongodb: inserting user")
		return user.User{}, err
	}

	newUser := &usr
	newUser.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return *newUser, nil
}

func (s *Store) Update(ctx context.Context, usr user.User) error {
	filter := bson.M{"_id": usr.ID}
	update := bson.M{"$set": usr}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
	}
	return err
}

func (s *Store) Delete(ctx context.Context, usr user.User) error {
	filter := bson.M{"_id": usr.ID}

	_, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
	}
	return err
}

func (s *Store) Query(ctx context.Context, filter user.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]user.User, error) {
	findOptions := options.Find()
	sortOrder := 1 // Default to ascending
	if orderBy.Direction == order.DESC {
		sortOrder = -1
	}
	findOptions.SetSort(bson.D{{Key: orderBy.Field, Value: sortOrder}})
	findOptions.SetSkip(int64((pageNumber - 1) * rowsPerPage))
	findOptions.SetLimit(int64(rowsPerPage))

	cur, err := s.collection.Find(ctx, filter, findOptions)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return nil, err
	}
	defer cur.Close(ctx)

	var users []user.User
	for cur.Next(ctx) {
		var usr user.User
		err := cur.Decode(&usr)
		if err != nil {
			s.log.Error().Err(err).Msg("mongodb")
			return nil, err
		}
		users = append(users, usr)
	}

	if err := cur.Err(); err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return nil, err
	}

	return users, nil
}

func (s *Store) Count(ctx context.Context, filter user.QueryFilter) (int, error) {
	count, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return 0, err
	}
	return int(count), nil
}

func (s *Store) QueryByID(ctx context.Context, userID uuid.UUID) (user.User, error) {
	filter := bson.M{"_id": userID}

	var usr user.User
	err := s.collection.FindOne(ctx, filter).Decode(&usr)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return user.User{}, err
	}
	return usr, nil
}

func (s *Store) QueryByIDs(ctx context.Context, userIDs []uuid.UUID) ([]user.User, error) {
	filter := bson.M{"_id": bson.M{"$in": userIDs}}

	cur, err := s.collection.Find(ctx, filter)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return nil, err
	}
	defer cur.Close(ctx)

	var users []user.User
	for cur.Next(ctx) {
		var usr user.User
		err := cur.Decode(&usr)
		if err != nil {
			s.log.Error().Err(err).Msg("mongodb")
			return nil, err
		}
		users = append(users, usr)
	}

	if err := cur.Err(); err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return nil, err
	}

	return users, nil
}

func (s *Store) QueryByEmail(ctx context.Context, email mail.Address) (user.User, error) {
	filter := bson.M{"email": email.Address}

	var usr user.User
	err := s.collection.FindOne(ctx, filter).Decode(&usr)
	if err != nil {
		s.log.Error().Err(err).Msg("mongodb")
		return user.User{}, err
	}
	return usr, nil
}
