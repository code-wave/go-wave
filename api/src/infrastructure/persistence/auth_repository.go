package persistence

import (
	"context"
	"log"
	"time"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/go-redis/redis/v8"
)

var _ repository.AuthRepository = &AuthRepo{}
var ctx = context.Background()

type AuthRepo struct {
	rClient *redis.Client
}

func NewAuthRepository(rClient *redis.Client) *AuthRepo {
	return &AuthRepo{
		rClient: rClient,
	}
}

func (ar *AuthRepo) Create(rt *entity.RefreshToken) *errors.RestErr {
	expUTC := time.Unix(rt.ExpiresAt, 0)
	if err := ar.rClient.Set(ctx, rt.Uuid, rt.RefreshToken, time.Until(expUTC)).Err(); err != nil {
		log.Println("error when save refresh token in redis")
		redisErr := errors.NewInternalServerError("redis error")
		return redisErr
	}

	return nil
}
func (ar *AuthRepo) Delete() {}
func (ar *AuthRepo) Fetch(uuid string) (string, *errors.RestErr) {
	rtRedis, err := ar.rClient.Get(ctx, uuid).Result()
	if err != nil {
		if err == redis.Nil {
			redisErr := errors.NewNotFoundError("not found value")
			return "", redisErr
		}
		redisErr := errors.NewUnauthorizedError("refresh token is not valid")
		log.Println("error when get refresh token in redis, ", err)
		return "", redisErr
	}
	return rtRedis, nil
}
