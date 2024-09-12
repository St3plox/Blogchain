package consumer

import (
	"context"
	"time"

	"github.com/St3plox/Blogchain/business/core/email"
	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/rs/zerolog"
)

// likeConsumer defines the interface for consuming Like events
type likeConsumer interface {
	Consume(ctx context.Context) (<-chan like.Like, error)
}

// Controller defines the structure that manages event consumption and email sending
type Controller struct {
	consumer    likeConsumer
	delay       time.Duration
	log         *zerolog.Logger
	emailSender email.EmailSender
}

// New creates a new Controller
func New(consumer likeConsumer, delay time.Duration, log *zerolog.Logger, emailSender email.EmailSender) *Controller {
	return &Controller{
		consumer:    consumer,
		delay:       delay,
		log:         log,
		emailSender: emailSender,
	}
}

// ListenForEvents listens to like events from Kafka and sends emails for each event
func (c *Controller) ListenForEvents(ctx context.Context) {
	for {
		// Consume events from the Kafka consumer
		likeEventChannel, err := c.consumer.Consume(ctx)
		if err != nil {
			c.log.Error().Err(err).Msg("Failed to start consuming events")
			return
		}

		for likeEvent := range likeEventChannel {
			c.log.Info().Msg("Started processing like event")

			err := c.emailSender.Send(ctx, email.LikeToEmail(likeEvent)) 
			if err != nil {
				c.log.Error().Err(err).Msg("Failed to send email after retries")
				continue
			}

			c.log.Info().Msg("Successfully sent email")
		}

		// time.Sleep(c.delay)
	}
}