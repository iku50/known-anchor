package service

import (
	"known-anchors/dal"
	"known-anchors/dal/db/dao"
	"known-anchors/dal/mq"
	"known-anchors/dal/redis"
	"sync"
)

type ServiceContext struct {
	DBQuery      *dao.Query
	Redis        *redis.RedisClient
	WorkProdecer *mq.Producer
}

var once sync.Once

func NewServiceContext() *ServiceContext {
	var sc ServiceContext
	once.Do(func() {
		sc.DBQuery = dao.Use(dal.DB.Debug())
		sc.Redis = redis.InitRedisClient()
		sc.WorkProdecer = mq.NewProducer()

		// init mq consumer
		c := mq.NewConsumer()
		go c.Consume()
	})
	return &sc
}
