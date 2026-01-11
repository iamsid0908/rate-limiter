package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

// RedisInit initializes the Redis client with proper context
func RedisInit() {
	client = redis.NewClient(&redis.Options{
		Addr:         GetConfig().RedisAddress,
		DB:           GetConfig().RedisDB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
	})

	// Test connection with context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status := client.Ping(ctx)
	_, err := status.Result()
	if err != nil {
		err = fmt.Errorf("redis connection error: %v", err)
		panic(err)
	}

	log.Println("✅ Connected to Redis successfully at", GetConfig().RedisAddress)
}

// RedisClient returns the Redis client instance
func RedisClient() *redis.Client {
	return client
}

// CloseRedis gracefully closes the Redis connection
func CloseRedis() error {
	if client != nil {
		log.Println("Closing Redis connection...")
		return client.Close()
	}
	return nil
}
