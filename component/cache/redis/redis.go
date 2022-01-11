package redis

import (
	"git.xios.club/xios/gc"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string `value:"${redis.addr}"`
	Password string `value:"${redis.password}"`
	DB       int    `value:"${redis.db}"`
}

func init() {
	gc.RegisterBeanFn(func(config RedisConfig) *redis.Client {
		return redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		})
	})
}
