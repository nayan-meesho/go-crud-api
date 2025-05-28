package models

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis(address string) {
	// Initialize Redis client
	RDB = redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})

	// Ping Redis to check if the connection is successful
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}