package consumer

import (
	"context"

	"github.com/St3plox/Blogchain/business/core/like"
)

type likeConsumer interface {
	Ingest(ctx context.Context) (<-chan like.Like, error)
}

type Controller struct {
	ingester likeConsumer
}

// New creates a rating service controller
func New(consumer likeConsumer) *Controller {
	return &Controller{consumer}
}

func ListenForEvents(	)
