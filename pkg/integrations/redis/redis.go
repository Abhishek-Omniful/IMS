package redisService

import (
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/redis"
)

var redisClient *redis.Client
var logger = log.DefaultLogger()

func ConnectRedis() {
	config := &redis.Config{
		Hosts:       []string{"127.0.0.1:6379"},
		PoolSize:    50,
		MinIdleConn: 10,
		DB:          0,
	}
	redisClient = redis.NewClient(config)
	logger.Info("Connecting to Redis...")
}

func GetRedis() *redis.Client {
	return redisClient
}
