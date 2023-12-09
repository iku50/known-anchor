package redis

import (
	"context"
	"known-anchors/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.DB,
	})
	return &RedisClient{
		Client: client,
	}
}

type RedisClientInterface interface {
	Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) error {
	return r.Client.Set(ctx, key, value, expireTime).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
