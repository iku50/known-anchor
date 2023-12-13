package mq

import (
	"context"
	"known-anchors/dal/redis"
)

type RedisMQ struct {
	redisClient *redis.RedisClient
}

func NewRedisMQ(redisClient *redis.RedisClient) *RedisMQ {
	return &RedisMQ{redisClient: redisClient}
}

func (r *RedisMQ) Produce(ctx context.Context, message *MQMessage) error {
	return r.redisClient.LPush(ctx, message.Topic, message.Body)
}

func (r *RedisMQ) Consume(ctx context.Context, topic string) (*MQMessage, error) {
	redisClient := r.redisClient
	result, err := redisClient.BRPop(ctx, topic, 0)
	if err != nil {
		return nil, err
	}
	return &MQMessage{
		Topic: topic,
		Body:  []byte(result[1]),
	}, nil
}
