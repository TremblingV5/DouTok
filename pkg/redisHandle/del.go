package redishandle

import "context"

func (c *RedisClient) Del(ctx context.Context, key string) error {
	return c.Client.Del(ctx, key).Err()
}
