package redishandle

import (
	"context"
	"encoding/json"
)

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}

func (c *RedisClient) GetObj(ctx context.Context, key string, out any) error {
	result, err := c.Get(ctx, key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), &out)
	if err != nil {
		return err
	}

	return nil
}
