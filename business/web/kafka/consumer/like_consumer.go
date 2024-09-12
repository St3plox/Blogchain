// Package consumer implements consumer logic for Apache Kafka
package consumer
//NOTE: will create more generic interface in future

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog"
)

// LikeConsumer defines a Kafka consumer
type LikeConsumer struct {
	consumer *kafka.Consumer
	topic    string
	log      *zerolog.Logger
}

var (
	ErrConsumer = fmt.Errorf("consumer error")
)

// NewLikeConsumer creates a new Kafka consumer
func NewLikeConsumer(adrr string, groupID string, topic string, log *zerolog.Logger) (*LikeConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": adrr,
		"group.id":          groupID,
	})
	if err != nil {
		return nil, err
	}
	return &LikeConsumer{consumer, topic, log}, nil
}

// Consume starts ingesting from Kafka and returns a channel of rating events
// representing the data consumed from topics
func (lc *LikeConsumer) Consume(ctx context.Context) (<-chan like.Like, error) {
	if err := lc.consumer.SubscribeTopics([]string{lc.topic}, nil); err != nil {
		return nil, err
	}

	ch := make(chan like.Like, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				lc.consumer.Close()
			default:
			}

			// Waiting indefinitely
			msg, err := lc.consumer.ReadMessage(-1)
			if err != nil {
				lc.log.Err(err)
				continue
			}
			var event like.Like
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				lc.log.Err(err)
				continue
			}
			ch <- event
		}
	}()

	return ch, nil
}
