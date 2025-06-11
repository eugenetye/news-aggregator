package redis

import (
	"github.com/redis/go-redis/v9"
	"context"
	"os"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:    os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB: 0, // use default DB
	})
	
}