package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaWriter(brokerAddress string, topic string) {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
}

func PublishMessage(msg string) {
	err := writer.WriteMessages(context.Background(), 
		kafka.Message{
			Value: []byte(msg),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message to kafka: %v", err)
	} else {
		log.Printf("Message published to kafka: %s", msg)
	}
}