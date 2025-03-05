package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	once        sync.Once
)

// Initialize Redis connection once
func ConnectToRedis() {
	once.Do(func() {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		password := os.Getenv("REDIS_PASSWORD")

		RedisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
			DB:       0,
		})

		ctx := context.Background()
		_, err := RedisClient.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}

		log.Println("Connected to Redis")
	})
}

func GetRedis() *redis.Client {
	if RedisClient == nil {
		log.Fatal("Redis is not initialized")
	}
	return RedisClient
}
