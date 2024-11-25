package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func main() {
	time.Sleep(time.Second*5)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "table_of_two_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	err = consumer.Subscribe("table_of_two", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Redis address
	})
	defer redisClient.Close()

	ctx := context.Background()

	fmt.Println("Listening for messages...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			key := string(msg.Key)
			value := string(msg.Value)

			// Insert into Redis
			err := redisClient.Set(ctx, key, value, 0).Err()
			if err != nil {
				log.Printf("Failed to insert into Redis: %s", err)
			} else {
				fmt.Printf("Inserted into Redis: %s -> %s\n", key, value)
			}
		} else {
			log.Printf("Consumer error: %v\n", err)
		}
	}
}
