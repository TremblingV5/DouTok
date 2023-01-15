package redishandle

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedis(dsn string, pwd string, dbs map[string]int) (map[string]*redis.Client, error) {
	redisCaches := make(map[string]*redis.Client)

	for k, v := range dbs {
		redisCaches[k] = redis.NewClient(&redis.Options{
			Addr:     dsn,
			Password: pwd,
			DB:       v,
			PoolSize: 20,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := redisCaches[k].Ping(ctx).Result()
		if err != nil {
			return redisCaches, err
		}
	}

	return redisCaches, nil
}
