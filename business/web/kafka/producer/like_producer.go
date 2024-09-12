// Package consumer implements producer logic for Apache Kafka
package producer
//NOTE: will create more generic interface in future

import (
	"encoding/json"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type LikeProducer struct {
	producer *kafka.Producer
}

func NewLikeProducer(producer *kafka.Producer) *LikeProducer {
	return &LikeProducer{producer}

}


func ProduceLikesEvents(producer *kafka.Producer, topic string, likesEvents []like.Like) error {
    for _, likeEvent := range likesEvents {
        encodedEvent, err := json.Marshal(likeEvent)
        if err != nil {
            return err
        }

        message := &kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          encodedEvent,
        }
        if err := producer.Produce(message, nil); err != nil {
            return err
        }
    }
    return nil
}
