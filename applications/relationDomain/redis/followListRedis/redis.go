package followListRedis

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/pkg/constants"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
	"github.com/go-redis/redis/v8"
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

func getRedisKeyByUserId(userId int64) string {
	return constants.FollowListPrefix + fmt.Sprintf("%d", userId)
}

func (c *Client) Get(ctx context.Context, userId int64) ([]int64, error) {
	key := getRedisKeyByUserId(userId)
	res, err := c.client.HGetAll(ctx, key)
	if err != nil {
		return nil, err
	}
	ret := make([]int64, 0)
	for k, v := range res {
		kI64, _ := strconv.ParseInt(k, 10, 64)
		if v == "1" {
			ret = append(ret, kI64)
		}
	}

	if len(ret) <= 0 {
		return ret, redis.Nil
	}
	return ret, nil
}

func (c *Client) Set(ctx context.Context, userId int64, relations []*model.Relation) error {
	val := make([]string, len(relations)*2)
	for _, v := range relations {
		val = append(val, fmt.Sprintf("%d", v.ToUserId))
		val = append(val, fmt.Sprintf("%d", v.Status))
	}

	key := getRedisKeyByUserId(userId)
	err := c.client.HSetMore(ctx, key, val...)
	if err != nil {
		return err
	}
	return nil
}
