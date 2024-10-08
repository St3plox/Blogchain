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

type MediaData struct {
	Filename string `json:"name"`
	Url      string `json:"url"`
}

// GenUrl generates part of url for downloading the mediafile
func (m Media) GenUrl() string {
	return "/v1/media/" + m.ID.Hex()
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

// GenerateMediaLists creates lists of media names and URLs from a slice of Media
func GenerateMediaLists(mediaList []Media) ([]string, []string) {
	var mediaNames []string
	var mediaUrls []string

	for _, media := range mediaList {
		mediaNames = append(mediaNames, media.Filename)
		mediaUrls = append(mediaUrls, media.GenUrl())
	}

	return mediaNames, mediaUrls
}

func MapTo(media Media) MediaData {
	return MediaData{
		Filename: media.Filename,
		Url:      media.GenUrl(),
	}
}

func MapToMultiple(media []Media) []MediaData {

	mediaData := make([]MediaData, 1)
	for _, m := range media {
		mediaData = append(mediaData, MapTo(m))
	}

	return mediaData
}
