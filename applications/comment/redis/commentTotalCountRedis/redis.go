package commentTotalCountRedis

import (
	"context"
	"fmt"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
)

type Client struct {
	client *redishandle.RedisClient
}

//go:generate mockgen -source=redis.go -destination=./mocks/redis_mock.go -package ComentTotalCountRedisRepositoryMocks
type IClient interface {
	Get(ctx context.Context, videoId int64) (int64, error)
	Delete(ctx context.Context, videoId ...int64) error
}

func NewClient(client *redishandle.RedisClient) *Client {
	return &Client{
		client: client,
	}
}

func getRedisKeyByVideoId(videoId int64) string {
	return fmt.Sprint("comment_total_count_", videoId)
}

func (c *Client) Get(ctx context.Context, videoId int64) (int64, error) {
	key := getRedisKeyByVideoId(videoId)
	return c.client.GetI64(ctx, key)
}

func (c *Client) Delete(ctx context.Context, videoId ...int64) error {
	videoIdList := make([]string, len(videoId))
	for i, id := range videoId {
		videoIdList[i] = getRedisKeyByVideoId(id)
	}
	return c.client.DelBatch(ctx, videoIdList...)
}
