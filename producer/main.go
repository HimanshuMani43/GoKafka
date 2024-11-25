package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	time.Sleep(time.Second*5)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	topic := "table_of_two"
	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("%dtimes2", i)
		value := fmt.Sprintf("%d", i*2)
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(value),
		}
		producer.Produce(message, nil)
		fmt.Printf("Produced message: %s -> %s\n", key, value)
	}

	// Wait for message delivery before exiting
	producer.Flush(15 * 1000)
}
