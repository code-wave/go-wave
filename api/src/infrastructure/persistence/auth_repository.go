package persistence

import (
	"context"
	"log"
	"strconv"
	"strings"
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

	//save redis[rt.Uuid] = (userID-rt)
	if err := ar.rClient.Set(ctx, rt.Uuid, strconv.FormatUint(rt.UserID, 10)+"-"+rt.RefreshToken, time.Until(expUTC)).Err(); err != nil {
		log.Println("error when save refresh token in redis")
		redisErr := errors.NewInternalServerError("redis error")
		return redisErr
	}
	return nil
}

func (ar *AuthRepo) Delete(uuid string) *errors.RestErr {
	deleted, err := ar.rClient.Del(ctx, uuid).Result()
	//del success, then return 1
	if err != nil || deleted != 1 {
		log.Println("error when delete refresh token in redis")
		redisErr := errors.NewUnauthorizedError("unauthorized, refresh token is not valid")
		return redisErr
	}
	return nil
}

func (ar *AuthRepo) Fetch(uuid string) (uint64, *errors.RestErr) {
	userIDAndRt, err := ar.rClient.Get(ctx, uuid).Result()
	if err != nil {
		if err == redis.Nil {
			redisErr := errors.NewUnauthorizedError("unauthorized, refresh token is expired please relogin")
			return 0, redisErr
		}
		redisErr := errors.NewUnauthorizedError("unauthorized, refresh token is expired please relogin")
		log.Println("error when get refresh token in redis, ", err)
		return 0, redisErr
	}

	//split userID + '-' + rt
	userID := strings.Split(userIDAndRt, "-")[0]
	uid, _ := strconv.ParseUint(userID, 10, 64)
	log.Println(userIDAndRt)
	log.Println(userID)
	log.Println(uid)
	return uid, nil
}
