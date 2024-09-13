// Package consumer implements consumer logic for Apache Kafka
package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/St3plox/Blogchain/business/web/broker"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
)

// LikeConsumer defines a Kafka consumer
type LikeConsumer struct {
	consumer      *kafka.Consumer
	topic         string
	log           *zerolog.Logger
	channelBuffer int // Buffer size for the event channel
	retryDelay    time.Duration
}

// ErrConsumer defines a consumer error
var ErrConsumer = fmt.Errorf("consumer error")

// NewLikeConsumer creates a new Kafka consumer
func NewLikeConsumer(adrr, groupID, topic string, log *zerolog.Logger, bufferSize int, retryDelay time.Duration) (*LikeConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": adrr,
		"group.id":          groupID,
		"auto.offset.reset": "earliest", // This can be configurable
	})
	if err != nil {
		return nil, err
	}

	// Return LikeConsumer instance with a configurable channel buffer and retry delay
	return &LikeConsumer{
		consumer:      consumer,
		topic:         topic,
		log:           log,
		channelBuffer: bufferSize,
		retryDelay:    retryDelay,
	}, nil
}

// Consume starts ingesting from Kafka and returns a channel of rating events
// representing the data consumed from topics
func (lc *LikeConsumer) Consume(ctx context.Context) (<-chan broker.LikeEvent, error) {

	if err := lc.consumer.SubscribeTopics([]string{lc.topic}, nil); err != nil {
		return nil, err
	}

	// Create a channel for emitting Like events
	ch := make(chan broker.LikeEvent, lc.channelBuffer)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		retryDelay := lc.retryDelay

		for {
			select {
			case <-ctx.Done():

				lc.log.Info().Msg("Shutting down consumer...")
				lc.consumer.Close()
				close(ch)
				return

			default:

				msg, err := lc.consumer.ReadMessage(-1)
				if err != nil {

					// Log the error and retry after a backoff period
					lc.log.Error().Err(err).Msg("Error reading message, retrying with backoff")
					time.Sleep(retryDelay)
					retryDelay *= 2
					if retryDelay > 30*time.Second {
						retryDelay = 30 * time.Second
					}
					continue
				}

				// Reset backoff after a successful read
				retryDelay = lc.retryDelay

				// Parse the Kafka message
				var event broker.LikeEvent
				if err := json.Unmarshal(msg.Value, &event); err != nil {
					lc.log.Error().Err(err).Msg("Error unmarshalling Kafka message")
					continue
				}

				// Send the event to the channel (non-blocking if channel is full)
				select {
				case ch <- event:
				case <-ctx.Done(): // In case the context is cancelled while sending
					lc.log.Info().Msg("Context cancelled while sending event")
					return
				}
			}
		}
	}()

	go func() {
		<-ctx.Done() // Wait for cancellation
		wg.Wait()    // Wait for the consumer goroutine to finish
		lc.log.Info().Msg("Consumer shutdown complete")
	}()

	return ch, nil
}
