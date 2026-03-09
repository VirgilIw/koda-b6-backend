package config

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type configRedis struct {
	rdbhost     string
	rdbport     string
	rdbpassword string
}

func InitRedis() *redis.Client {
	cfg := configRedis{
		rdbhost:     os.Getenv("REDIS_HOST"),
		rdbport:     os.Getenv("REDIS_PORT"),
		rdbpassword: os.Getenv("REDIS_PASSWORD"),
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.rdbhost, cfg.rdbport),
		Password: cfg.rdbpassword,
		DB:       0,
	})

	return rdb
}
