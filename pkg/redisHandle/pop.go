package redishandle

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func (c *RedisClient) LPop(ctx context.Context, key string) (string, error) {
	return c.Client.LPop(ctx, key).Result()
}

func (c *RedisClient) LPops(ctx context.Context, key string, times int) ([]string, error) {
	l, err := c.ListSize(ctx, key)
	if err != nil {
		return nil, err
	}
	if l < int64(times) {
		return nil, ErrNotEnoughOpNumsInList
	}

	p := c.Client.TxPipeline()

	cmds := []*redis.StringCmd{}

	for i := 0; i < times; i++ {
		cmds = append(cmds, p.LPop(ctx, key))
	}

	_, err = p.Exec(ctx)

	if err != nil {
		return nil, err
	}

	result := []string{}

	for _, cmd := range cmds {
		result = append(result, cmd.Val())
	}

	return result, nil
}
