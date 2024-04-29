package database

import (
	"forest-run/realtime-runs-server/boot"
	"github.com/redis/go-redis/v9"
)

type DBAdapter interface {
}

type RedisAdapter struct {
	redis *redis.Client
}

func NewRedisAdapter(config boot.RedisConfig) DBAdapter {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})

	return &RedisAdapter{redis: client}
}
