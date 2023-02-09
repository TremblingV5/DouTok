package redishandle

import "context"

func (c *RedisClient) DelKey(ctx context.Context, keys ...string) error {
	return c.Client.Del(ctx, keys...).Err()
}
