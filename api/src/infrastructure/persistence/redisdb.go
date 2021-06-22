package persistence

import (
	"context"
	"log"

	"github.com/code-wave/go-wave/domain/repository"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Auth    repository.AuthRepository
	rClient *redis.Client
}

func NewRedisDB(host, port, password string) (*RedisService, error) {
	rClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	if err := rClient.Ping(context.Background()).Err(); err != nil {
		log.Println("redis server doesn't connect")
		return nil, err
	}

	log.Println("redis connected successfully")

	return &RedisService{
		Auth:    NewAuthRepository(rClient),
		rClient: rClient,
	}, nil
}
