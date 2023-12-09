package service

import (
	"context"
	"known-anchors/dal"
	"known-anchors/dal/db/dao"
	"known-anchors/dal/redis"
	"sync"
)

type ServiceContext struct {
	DBQuery *dao.Query
	Redis   *redis.RedisClientInterface
	Ctx     context.Context
}

var once sync.Once

func NewServiceContext() *ServiceContext {
	var sc ServiceContext
	once.Do(func() {
		sc.DBQuery = dao.Use(dal.DB.Debug())
		var rc redis.RedisClientInterface = redis.NewRedisClient()
		sc.Redis = &rc
		sc.Ctx = context.Background()
	})
	return &sc
}
