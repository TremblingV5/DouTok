package redishandle

import (
	"context"
	"encoding/json"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (c *RedisClient) SGet(ctx context.Context, key string) ([]string, error) {
	return c.Client.SMembers(ctx, key).Result()
}

func (c *RedisClient) SGetObj(ctx context.Context, key string) ([]*user.User, error) {
	res, err := c.SGet(ctx, key)
	if err != nil {
		return nil, err
	}
	result := make([]*user.User, len(res))
	for i, v := range res {
		t := &user.User{}
		err := json.Unmarshal([]byte(v), t)
		result[i] = t
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
