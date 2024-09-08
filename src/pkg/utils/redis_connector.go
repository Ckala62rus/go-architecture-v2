package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	RedisDb *RedisCache
)

func init() {
	host := MainConfig.RedisConfig.Host
	port := MainConfig.RedisConfig.Port
	db := MainConfig.RedisConfig.Db

	RedisDb = NewRedisCache(host+":"+port, db)

	// check connection stats
	ctx := context.Background()
	_, err := RedisDb.getClient().Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("***************** REDIS RUN *****************")
}

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
	defer func(rbd *redis.Client) {
		err := rbd.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rbd)
	res := rbd.Set(ctx, "user:"+token, token, time.Second*30)
	fmt.Println(res)
	return res
}

func (cache RedisCache) DeleteToken(
	ctx context.Context,
	token string,
) error {
	rbd := cache.getClient()
	iter := rbd.Scan(ctx, 0, "user:"+token, 0).Iterator()
	for iter.Next(ctx) {
		rbd.Del(ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (cache RedisCache) GetToken(
	ctx context.Context,
	token string,
) (string, error) {
	rbd := RedisDb.getClient()
	res := rbd.Get(ctx, "user:"+token)

	err := res.Err()
	if err != nil {
		return "", err
	}

	tokenRedis := res.Val()
	return tokenRedis, nil
}
