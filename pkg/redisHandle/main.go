package redishandle

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func InitRedis(dsn string, pwd string, dbs map[string]int) (map[string]*RedisClient, error) {
	redisCaches := make(map[string]*RedisClient)

	for k, v := range dbs {
		redisCaches[k] = &RedisClient{
			Client: redis.NewClient(&redis.Options{
				Addr:     dsn,
				Password: pwd,
				DB:       v,
				PoolSize: 20,
			}),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := redisCaches[k].Client.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}
	}

	return redisCaches, nil
}

func InitRedisClient(dsn string, pwd string, database int) (*redis.Client, error) {
	Client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: pwd,
		DB:       database,
		PoolSize: 20,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return Client, nil
}
