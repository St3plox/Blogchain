// Package consumer implements producer logic for Apache Kafka
package producer

//NOTE: will create more generic interface in future

import (
	"encoding/json"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = "likes"

type Producer interface {
	ProduceLikesEvents(likesEvents []like.LikeEvent) error
}

type LikeProducer struct {
	producer *kafka.Producer
}

func NewLikeProducer(producer *kafka.Producer) *LikeProducer {
	return &LikeProducer{producer}

}

func (lp *LikeProducer) ProduceLikesEvents(likesEvents []like.LikeEvent) error {
	for _, likeEvent := range likesEvents {
		encodedEvent, err := json.Marshal(likeEvent)
		if err != nil {
			return err
		}

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          encodedEvent,
		}
		if err := lp.producer.Produce(message, nil); err != nil {
			return err
		}
	}
	return nil
}
