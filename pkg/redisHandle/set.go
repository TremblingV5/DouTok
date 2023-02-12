package redishandle

import (
	"context"
	"encoding/json"
	"time"
)

func (c *RedisClient) Set(ctx context.Context, key string, value string, expire time.Duration) error {
	return c.Client.Set(ctx, key, value, expire).Err()
}

func (c *RedisClient) SetObj(ctx context.Context, key string, value any, expire time.Duration) error {
	valueJson, err := json.Marshal(value)

	if err != nil {
		return err
	}

	if err := c.Set(ctx, key, string(valueJson), expire); err != nil {
		return err
	}

	return nil
}

func (c *RedisClient) HSet(ctx context.Context, key string, hKey string, hValue string) error {
	return c.Client.HSet(ctx, key, hKey, hValue).Err()
}

func (c *RedisClient) HSetMore(ctx context.Context, key string, values ...string) error {
	return c.Client.HSet(ctx, key, values).Err()
}
