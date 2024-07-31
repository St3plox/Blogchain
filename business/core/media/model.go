package media

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerID   primitive.ObjectID `json:"owner_id,omitempty"`
	Filename  string             `json:"name"`
	Length    int64              `json:"length"`
	FileBytes []byte             `json:"fileBytes"`
}

type NewMedia struct {
	Filename  string `json:"name"`
	Length    int64  `json:"length"`
	FileBytes []byte `json:"fileBytes"`
}

// GenUrl generates part of url for downloading the mediafile
func (m Media) GenUrl() string {
	return "/media/" + m.ID.Hex()
}

func (m Media) CacheKey() string {
	return IdToCacheKey(m.ID.Hex())
}

func (p Media) CacheExpiration() time.Duration {
	return 2 * time.Hour
}

func IdToCacheKey(idHex string) string {
	return "media:" + idHex
}
