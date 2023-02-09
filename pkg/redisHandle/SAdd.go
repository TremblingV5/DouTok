package redishandle

import (
	"context"
	"encoding/json"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (c *RedisClient) SAdd(ctx context.Context, key string, value ...string) error {
	return c.Client.SAdd(ctx, key, value).Err()
}

func (c *RedisClient) SAddObj(ctx context.Context, key string, value ...*user.User) error {
	for _, v := range value {
		jsonValue, err := json.Marshal(v)
		if err != nil {
			return err
		}
		if err := c.SAdd(ctx, key, string(jsonValue)); err != nil {
			return err
		}
	}
	return nil
}
