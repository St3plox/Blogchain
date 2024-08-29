package likedb_test

import (
	"context"
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

	t.Run("QueryAllByPostID", func(t *testing.T) {
		postID := primitive.NewObjectID()
		like1 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     primitive.NewObjectID(),
			PostID:     postID,
			IsPositive: true,
		}
		like2 := like.Like{
			ID:         primitive.NewObjectID(),
			UserID:     primitive.NewObjectID(),
			PostID:     postID,
			IsPositive: false,
		}

		// Insert likes
		_, err := store.Create(ctx, like1)
		require.NoError(t, err)
		_, err = store.Create(ctx, like2)
		require.NoError(t, err)

		// Test QueryAllByPostID
		likes, err := store.QueryAllByPostID(ctx, postID.Hex())
		require.NoError(t, err)
		assert.Len(t, likes, 2)
		assert.Equal(t, postID, likes[0].PostID)
		assert.Equal(t, postID, likes[1].PostID)
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
