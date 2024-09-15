package likedb_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/St3plox/Blogchain/business/core/like/likedb"
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

func TestStore(t *testing.T) {
	ctx := context.Background()
	store := likedb.NewStore(&logger, testEnv.MongoClient)

	t.Run("Create and QueryByID", func(t *testing.T) {
		likeID := primitive.NewObjectID()
		newLike := like.Like{
			ID:         likeID,
			UserID:     primitive.NewObjectID(),
			PostID:     1,
			IsPositive: true,
		}

		// Test Create
		createdLike, err := store.Create(ctx, newLike)
		require.NoError(t, err)
		assert.NotNil(t, createdLike)

		// Test QueryByID
		resultLike, err := store.QueryByID(ctx, createdLike.ID.Hex())
		require.NoError(t, err)
		assert.Equal(t, createdLike, resultLike)
	})

	t.Run("QueryAllByUserID", func(t *testing.T) {
		userID := primitive.NewObjectID()
		like1 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     userID,
			PostID:     int64(1),
			IsPositive: true,
		}
		like2 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     userID,
			PostID:     int64(1),
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

	t.Run("QueryAllByPostID", func(t *testing.T) {

		like1 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     primitive.NewObjectID(),
			PostID:     int64(1),
			IsPositive: true,
		}
		like2 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     primitive.NewObjectID(),
			PostID:     int64(1),
			IsPositive: false,
		}

		// Insert likes
		_, err := store.Create(ctx, like1)
		require.NoError(t, err)
		_, err = store.Create(ctx, like2)
		require.NoError(t, err)

		// Test QueryAllByPostID
		likes, err := store.QueryAllByPostID(ctx, fmt.Sprint(like1.PostID))

		require.NoError(t, err)
		assert.Equal(t, like1.PostID, likes[0].PostID)
		assert.Equal(t, like1.PostID, likes[1].PostID)
	})

	t.Run("Update and DeleteByID", func(t *testing.T) {
		likeID := primitive.NewObjectID()
		initialLike := like.Like{
			ID:         likeID,
			UserID:     primitive.NewObjectID(),
			PostID:     int64(1),
			IsPositive: true,
		}

		// Create initial like
		createdLike, err := store.Create(ctx, initialLike)
		require.NoError(t, err)

		// Test Update
		updatedLike := like.Like{
			ID:         createdLike.ID,
			UserID:     initialLike.UserID,
			PostID:     initialLike.PostID,
			IsPositive: false,
		}
		resultLike, err := store.Update(ctx, updatedLike)
		require.NoError(t, err)
		assert.Equal(t, updatedLike, resultLike)

		// Test DeleteByID
		err = store.DeleteByID(ctx, resultLike.ID.Hex())
		require.NoError(t, err)

		// Verify deletion
		_, err = store.QueryByID(ctx, likeID.Hex())
		assert.True(t, like.IsLikeNotFound(err))
	})

	t.Run("QueryByUserAndPostID - Success", func(t *testing.T) {
		// Create a like for testing
		userID := primitive.NewObjectID()
		postID := int64(1)
		newLike := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     userID,
			PostID:     postID,
			IsPositive: true,
		}

		// Insert the like
		insertedLike, err := store.Create(ctx, newLike)
		require.NoError(t, err)

		// Test QueryByUserAndPostID
		resultLike, err := store.QueryByUserAndPostID(ctx, userID.Hex(), postID)
		require.NoError(t, err)
		assert.Equal(t, insertedLike.ID, resultLike.ID)
		assert.Equal(t, insertedLike.UserID, resultLike.UserID)
		assert.Equal(t, insertedLike.PostID, resultLike.PostID)
		assert.Equal(t, insertedLike.IsPositive, resultLike.IsPositive)
	})

	t.Run("QueryByUserAndPostID - Invalid UserID", func(t *testing.T) {
		// Test with an invalid user ID
		invalidUserID := "invalidUserID"
		postID := int64(1)

		_, err := store.QueryByUserAndPostID(ctx, invalidUserID, postID)
		assert.Error(t, err)
		assert.True(t, like.IsInvalidUserIDFormat(err))
	})

	t.Run("QueryByUserAndPostID - Like Not Found", func(t *testing.T) {
		// Test with a valid user ID but no corresponding like
		userID := primitive.NewObjectID().Hex()
		postID := int64(99) // A post ID that doesn't exist

		_, err := store.QueryByUserAndPostID(ctx, userID, postID)
		assert.Error(t, err)
		assert.True(t, like.IsLikeNotFound(err))
	})
}
