package service

import (
	"context"
	"errors"
	"time"

	"github.com/TremblingV5/DouTok/pkg/constants"
)

/*
	从Redis中获取缓存的feed列表，通过Redis事务执行若干次feed操作，从而获得足够的feed list
*/
func GetFeedCache(ctx context.Context, user_id string, num int64) ([]VideoInHB, bool) {
	res, err := RedisClients[constants.FeedSendBox].LPops(ctx, user_id, int(num))
	if err != nil {
		return []VideoInHB{}, false
	}

	video_list := String2VideoList(res)

	if err != nil {
		return []VideoInHB{}, false
	}

	return video_list, true
}

/*
	将新的feed列表存储到Redis中，method参数只允许l或r，代表选择不同的方法Push到Redis
*/
func SetFeedCache(ctx context.Context, method string, userId string, values ...VideoInHB) error {
	videoList := VideoList2String(values)
	switch method {
	case "l":
		return RedisClients[constants.FeedSendBox].LPush(ctx, userId, videoList...)
	case "r":
		return RedisClients[constants.FeedSendBox].RPush(ctx, userId, videoList...)
	default:
		return errors.New("unknown method, only accept 'l' or 'r'")
	}
}

/*
	获取某个user_id在系统中的marked_time
*/
func GetMarkedTime(ctx context.Context, user_id string) (string, error) {
	return RedisClients[constants.TimeCache].Get(ctx, user_id)
}

/*
	为某个user_id设置新的marked_time
*/
func SetMarkedTime(ctx context.Context, user_id string, marked_time string) error {
	return RedisClients[constants.TimeCache].Set(ctx, user_id, marked_time, 24*time.Hour)
}
