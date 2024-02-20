package favorite_redis

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

func getKey(userId, videoId int64) string {
	return fmt.Sprint("favorite_userId_", userId, "_videoId_", videoId)
}

func (r *Redis) Load(ctx context.Context, userId, videoId int64) (bool, error) {
	return r.client.Get(ctx, getKey(userId, videoId)).Bool()
}

func (r *Redis) Put(ctx context.Context, userId, videoId int64, ttl time.Duration) error {
	return r.client.Set(ctx, getKey(userId, videoId), 1, ttl).Err()
}

func (r *Redis) Remove(ctx context.Context, userId, videoId int64, ttl time.Duration) error {
	return r.client.Set(ctx, getKey(userId, videoId), 0, ttl).Err()
}
