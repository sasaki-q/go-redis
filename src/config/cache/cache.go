package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Data struct {
	K string
	V string
}

var (
	c       context.Context
	myCache *redis.Client
)

func Init() {
	c = context.Background()

	myCache = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func Set(d Data) {
	err := myCache.Set(c, d.K, d.V, 5*time.Second).Err()

	if err != nil {
		panic(err)
	}
}

func Read(key string) string {
	val, err := myCache.Get(c, key).Result()

	if err != nil || val == "" {
		return ""
	}

	return val
}
