package redishandle

import "context"

func (c *RedisClient) LPush(ctx context.Context, key string, value string) error {
	return c.Client.LPush(ctx, key, value).Err()
}

func (c *RedisClient) RPush(ctx context.Context, key string, value string) error {
	return c.Client.RPush(ctx, key, value).Err()
}
