package persistence

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRedisClient(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Println("unable to connect to redis", err.Error())
		return
	}

	log.Println("redis connected successfully")
}
