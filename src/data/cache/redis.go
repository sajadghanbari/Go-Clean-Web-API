package cache

import (
	"clean-web-api/config"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config){
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",cfg.Redis.Host,cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB: 0,
		DialTimeout: cfg.Redis.DialTimeout * time.Second,
		ReadTimeout: cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
		PoolSize: cfg.Redis.PoolSize,
		IdleTimeout: 500 * time.Millisecond,
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Microsecond,
	})
}

func GetRedis() *redis.Client{
	return redisClient
}

func CloseRedis(){
	redisClient.Close()
}