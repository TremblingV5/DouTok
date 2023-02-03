package redishandle

import (
	"context"
	"encoding/json"
)

func (c *RedisClient) SGet(ctx context.Context, key string) ([]string, error) {
	return c.Client.SMembers(ctx, key).Result()
}

func (c *RedisClient) SGetObj(ctx context.Context, key string) ([]any, error) {
	res, err := c.SGet(ctx, key)
	if err != nil {
		return nil, err
	}
	result := make([]any, len(res))
	for i, v := range res {
		err := json.Unmarshal([]byte(v), result[i])
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
