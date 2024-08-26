package mediadb_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/St3plox/Blogchain/business/core/media/mediadb"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string
var mongoClient *mongo.Client

// TestMain sets up the MongoDB container and runs the tests
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

// Function to start the MongoDB container using TestContainers
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

// Example test for creating media in the MongoDB database
func TestStore_Create(t *testing.T) {
	ctx := context.Background()

	// Set up logger and create the Store instance
	logger := zerolog.New(os.Stdout)
	store := mediadb.NewStore(&logger, mongoClient)

	// Test data
	media := media.Media{
		OwnerID:   primitive.NewObjectID(),
		Filename:  "testfile.jpg",
		Length:    12345,
		FileBytes: []byte{1, 2, 3},
	}

	// Insert media into MongoDB
	result, err := store.Create(ctx, media)
	require.NoError(t, err)
	assert.NotEmpty(t, result.ID)

	// Retrieve media from MongoDB
	foundMedia, err := store.QueryByID(ctx, result.ID.Hex())
	require.NoError(t, err)
	assert.Equal(t, result.ID, foundMedia.ID)
	assert.Equal(t, "testfile.jpg", foundMedia.Filename)
}

// Example test for deleting media by ID in the MongoDB database
func TestStore_DeleteByID(t *testing.T) {
	ctx := context.Background()

	// Set up logger and create the Store instance
	logger := zerolog.New(os.Stdout)
	store := mediadb.NewStore(&logger, mongoClient)

	// Test data
	testMedia := media.Media{
		OwnerID:   primitive.NewObjectID(),
		Filename:  "todeletefile.jpg",
		Length:    12345,
		FileBytes: []byte{1, 2, 3},
	}

	// Insert media into MongoDB
	result, err := store.Create(ctx, testMedia)
	require.NoError(t, err)
	assert.NotEmpty(t, result.ID)

	// Delete the media by ID
	err = store.DeleteByID(ctx, result.ID.Hex())
	require.NoError(t, err)

	// Verify that the media no longer exists
	_, err = store.QueryByID(ctx, result.ID.Hex())
	require.Error(t, err)
	assert.Equal(t, media.ErrNotFound, err)
}

func TestStore_CreateMultiple(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	store := mediadb.NewStore(&logger, mongoClient)

	mediaList := []media.Media{
		{
			OwnerID:   primitive.NewObjectID(),
			Filename:  "file1.jpg",
			Length:    12345,
			FileBytes: []byte{1, 2, 3},
		},
		{
			OwnerID:   primitive.NewObjectID(),
			Filename:  "file2.jpg",
			Length:    54321,
			FileBytes: []byte{4, 5, 6},
		},
	}

	result, err := store.CreateMultiple(ctx, mediaList)
	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.NotEmpty(t, result[0].ID)
	assert.NotEmpty(t, result[1].ID)
}

func TestStore_Update(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	store := mediadb.NewStore(&logger, mongoClient)

	media := media.Media{
		OwnerID:   primitive.NewObjectID(),
		Filename:  "file_to_update.jpg",
		Length:    12345,
		FileBytes: []byte{1, 2, 3},
	}

	createdMedia, err := store.Create(ctx, media)
	require.NoError(t, err)

	createdMedia.Filename = "updated_file.jpg"
	updatedMedia, err := store.Update(ctx, createdMedia)
	require.NoError(t, err)

	assert.Equal(t, "updated_file.jpg", updatedMedia.Filename)

	foundMedia, err := store.QueryByID(ctx, updatedMedia.ID.Hex())
	require.NoError(t, err)
	assert.Equal(t, "updated_file.jpg", foundMedia.Filename)
}

func TestStore_Delete(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	store := mediadb.NewStore(&logger, mongoClient)

	testMedia := media.Media{
		OwnerID:   primitive.NewObjectID(),
		Filename:  "file_to_delete.jpg",
		Length:    12345,
		FileBytes: []byte{1, 2, 3},
	}

	createdMedia, err := store.Create(ctx, testMedia)
	require.NoError(t, err)

	err = store.Delete(ctx, createdMedia)
	require.NoError(t, err)

	_, err = store.QueryByID(ctx, createdMedia.ID.Hex())
	assert.Error(t, err)
	assert.Equal(t, media.ErrNotFound, err)
}

func TestStore_QueryByIDs(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(os.Stdout)
	store := mediadb.NewStore(&logger, mongoClient)

	mediaList := []media.Media{
		{
			OwnerID:   primitive.NewObjectID(),
			Filename:  "file1.jpg",
			Length:    12345,
			FileBytes: []byte{1, 2, 3},
		},
		{
			OwnerID:   primitive.NewObjectID(),
			Filename:  "file2.jpg",
			Length:    54321,
			FileBytes: []byte{4, 5, 6},
		},
	}

	createdMediaList, err := store.CreateMultiple(ctx, mediaList)
	require.NoError(t, err)

	ids := []string{
		createdMediaList[0].ID.Hex(),
		createdMediaList[1].ID.Hex(),
	}

	foundMediaList, err := store.QueryByIDs(ctx, ids)
	require.NoError(t, err)
	assert.Len(t, foundMediaList, 2)
	assert.Equal(t, createdMediaList[0].ID, foundMediaList[0].ID)
	assert.Equal(t, createdMediaList[1].ID, foundMediaList[1].ID)
}
