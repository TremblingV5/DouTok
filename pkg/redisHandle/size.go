package redishandle

import "context"

func (c *RedisClient) ListSize(ctx context.Context, key string) (int64, error) {
	l, err := c.Client.LLen(ctx, key).Result()

	if err != nil {
		return -1, err
	}

	return l, nil
}
