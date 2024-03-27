package followCountRedis

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/constants"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"strconv"
)

type Client struct {
	client *redishandle.RedisClient
}

func NewClient(client *redishandle.RedisClient) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Get(ctx context.Context, userId int64) (int64, error) {
	ret, err := c.client.HGet(ctx, fmt.Sprintf("%d", userId), constants.FollowCount)
	if err != nil {
		return 0, err
	}
	follow, err := strconv.ParseInt(ret, 10, 64)
	if err != nil {
		return 0, err
	}
	return follow, nil
}

func (c *Client) Set(ctx context.Context, userId, count int64) error {
	err := c.client.HSet(ctx, fmt.Sprintf("%d", userId), constants.FollowCount, fmt.Sprintf("%d", count))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Del(ctx context.Context, userId int64) error {
	_, err := c.client.HDel(ctx, fmt.Sprintf("%d", userId), constants.FollowCount)
	if err != nil {
		return err
	}
	return nil
}
