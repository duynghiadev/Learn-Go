package datastore

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(ctx context.Context, host string, port string, password string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	// Test the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
		return nil, err
	}

	log.Println("Redis connection established successfully")
	return rdb, nil
}
