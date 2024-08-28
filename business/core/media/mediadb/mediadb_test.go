package mediadb_test

import (
	"context"
	"os"
	"testing"

	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/St3plox/Blogchain/business/core/media/mediadb"
	"github.com/St3plox/Blogchain/foundation/web/testutil"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var testEnv *testutil.TestEnv
var logger zerolog.Logger

func TestMain(m *testing.M) {
	ctx := context.Background()

	var err error
	testEnv, err = testutil.SetupMongoDBContainer(ctx)
	if err != nil {
		panic(err)
	}

	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	code := m.Run()

	if err := testEnv.Teardown(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(code)
}

func TestStore_Create(t *testing.T) {
	ctx := context.Background()

	store := mediadb.NewStore(&logger, testEnv.MongoClient)

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

	store := mediadb.NewStore(&logger, testEnv.MongoClient)

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

	store := mediadb.NewStore(&logger, testEnv.MongoClient)

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

	store := mediadb.NewStore(&logger, testEnv.MongoClient)

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

	store := mediadb.NewStore(&logger, testEnv.MongoClient)

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

	store := mediadb.NewStore(&logger, testEnv.MongoClient)

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
