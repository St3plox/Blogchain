package like

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	PostID     int64              `json:"post_id" bson:"post_id,omitempty"`
	IsPositive bool               `json:"is_positive"`
}

type NewLike struct {
	PostID     int64 `json:"post_id"`
	IsPositive bool  `json:"is_positive"`
}

// LikeEvent struct is used send events to Notification service
type LikeEvent struct {
	Like
	UserEmail string `json:"user_email"`
}

func (l Like) CacheKey() string {
	return IdToCacheKey(l.ID.Hex())
}

func (p Like) CacheExpiration() time.Duration {
	return 2 * time.Hour
}

func IdToCacheKey(idHex string) string {
	return "like:" + idHex
}
