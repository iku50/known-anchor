package service

import (
	"known-anchors/dal"
	"known-anchors/dal/db/dao"
	"known-anchors/dal/redis"
	"sync"
)

type ServiceContext struct {
	DBQuery *dao.Query
	Redis   *redis.RedisClient
}

var once sync.Once

func NewServiceContext() *ServiceContext {
	var sc ServiceContext
	once.Do(func() {
		sc.DBQuery = dao.Use(dal.DB.Debug())
		sc.Redis = redis.InitRedisClient()
	})
	return &sc
}
