package favorite_count_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	client *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func getRedisKey(videoId int64) string {
	return fmt.Sprint("favorite_count_", videoId)
}

func (r *Redis) Get(ctx context.Context, videoId int64) (int64, bool) {
	value, err := r.client.Get(ctx, getRedisKey(videoId)).Int64()
	if err != nil {
		return 0, false
	}

	return value, true
}

func (r *Redis) Set(ctx context.Context, videoId, count int64, ttl time.Duration) error {
	return r.client.Set(ctx, getRedisKey(videoId), count, ttl).Err()
}

func (r *Redis) Clear(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}
