package redis

import (
	"context"
	"known-anchors/config"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

var once sync.Once
var redisClient *RedisClient

func InitRedisClient() *RedisClient {
	once.Do(func() {
		redisClient = &RedisClient{
			Client: redis.NewClient(&redis.Options{
				Addr:     config.Conf.Redis.Addr,
				Password: config.Conf.Redis.Password,
				DB:       config.Conf.Redis.DB,
			}),
		}
	})
	return redisClient
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
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

func (r *RedisClient) LPush(ctx context.Context, key string, value interface{}) error {
	return r.Client.LPush(ctx, key, value).Err()
}

func (r *RedisClient) BRPop(ctx context.Context, key string, timeout time.Duration) ([]string, error) {
	return r.Client.BRPop(ctx, timeout, key).Result()
}
