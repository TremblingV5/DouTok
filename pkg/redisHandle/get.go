package redishandle

import (
	"context"
	"encoding/json"
)

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}

func (c *RedisClient) GetObj(ctx context.Context, key string, out any) (any, error) {
	result, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(result), out); err != nil {
		return nil, err
	}

	return out, nil
}
