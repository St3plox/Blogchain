package likedb_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/St3plox/Blogchain/business/core/like/likedb"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoURI    string
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Start MongoDB container
	mongoC, err := startMongoContainer(ctx)
	if err != nil {
		panic(err)
	}
	defer mongoC.Terminate(ctx)

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// Ensure MongoDB connection is established
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	mongoClient = client

	// Run the tests
	code := m.Run()

	// Disconnect MongoDB after tests are done
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func startMongoContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(10 * time.Second),
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	// Get MongoDB URI
	host, err := mongoC.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		return nil, err
	}

	mongoURI = "mongodb://" + host + ":" + port.Port()
	return mongoC, nil
}

func TestStore(t *testing.T) {
	ctx := context.Background()
	log := zerolog.New(os.Stdout) // Use logger or configure as needed
	store := likedb.NewStore(&log, mongoClient)

	t.Run("Create and QueryByID", func(t *testing.T) {
		likeID := primitive.NewObjectID()
		newLike := like.Like{
			ID:         likeID,
			UserID:     primitive.NewObjectID(),
			PostID:     primitive.NewObjectID(),
			IsPositive: true,
		}

		// Test Create
		createdLike, err := store.Create(ctx, newLike)
		require.NoError(t, err)
		assert.Equal(t, newLike, createdLike)

		// Test QueryByID
		resultLike, err := store.QueryByID(ctx, likeID.Hex())
		require.NoError(t, err)
		assert.Equal(t, newLike, resultLike)
	})

	t.Run("QueryAllByUserID", func(t *testing.T) {
		userID := primitive.NewObjectID()
		like1 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     userID,
			PostID:     primitive.NewObjectID(),
			IsPositive: true,
		}
		like2 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     userID,
			PostID:     primitive.NewObjectID(),
			IsPositive: false,
		}

		// Insert likes
		_, err := store.Create(ctx, like1)
		require.NoError(t, err)
		_, err = store.Create(ctx, like2)
		require.NoError(t, err)

		// Test QueryAllByUserID
		likes, err := store.QueryAllByUserID(ctx, userID.Hex())
		require.NoError(t, err)
		assert.Len(t, likes, 2)
	})

	t.Run("Update and DeleteByID", func(t *testing.T) {
		likeID := primitive.NewObjectID()
		initialLike := like.Like{
			ID:         likeID,
			UserID:     primitive.NewObjectID(),
			PostID:     primitive.NewObjectID(),
			IsPositive: true,
		}

		// Create initial like
		_, err := store.Create(ctx, initialLike)
		require.NoError(t, err)

		// Test Update
		updatedLike := like.Like{
			ID:         likeID,
			UserID:     initialLike.UserID,
			PostID:     initialLike.PostID,
			IsPositive: false,
		}
		resultLike, err := store.Update(ctx, updatedLike)
		require.NoError(t, err)
		assert.Equal(t, updatedLike, resultLike)

		// Test DeleteByID
		err = store.DeleteByID(ctx, likeID.Hex())
		require.NoError(t, err)

		// Verify deletion
		_, err = store.QueryByID(ctx, likeID.Hex())
		assert.True(t, likedb.IsLikeNotFound(err))
	})
}
