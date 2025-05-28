package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker, topic string) {
	r := kafka.NewReader(kafka.ReaderConfig {
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "task-group",
	})

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading kafka message: %v", err)
				continue
			}
			log.Printf("Notification received: %s", string(m.Value))
		}
	}()
}