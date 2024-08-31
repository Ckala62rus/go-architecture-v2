package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type JWTCache interface {
	SetToken(token string) *redis.StatusCmd
	GetToken(token string)
}

type RedisCache struct {
	host string
	db   int
}

func NewRedisCache(host string, db int) *RedisCache {
	return &RedisCache{
		host: host,
		db:   db,
	}
}

func (cache RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache RedisCache) SetToken(
	ctx context.Context,
	token string,
) *redis.StatusCmd {
	rbd := cache.getClient()
	res := rbd.Set(ctx, "user:1:secret_token", token, time.Second*10)
	return res
}
