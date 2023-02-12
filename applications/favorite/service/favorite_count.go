package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/go-redis/redis/v8"
)

func QueryFavoriteCount(video_id []int64) (map[int64]int64, error) {
	res, err := DoFavoriteCnt.Where(
		FavoriteCnt.VideoId.In(video_id...),
	).Find()

	if err != nil {
		return nil, err
	}

	resMap := make(map[int64]int64)
	for _, v := range res {
		resMap[v.VideoId] = v.Number
	}

	return resMap, nil
}

func AddCount(video_id int64) error {
	_, err := DoFavoriteCnt.Where(
		FavoriteCnt.VideoId.Eq(video_id),
	).Update(FavoriteCnt.Number, FavoriteCnt.Number.Add(1))
	return err
}

func ReduceCount(video_id int64) error {
	_, err := DoFavoriteCnt.Where(
		FavoriteCnt.VideoId.Eq(video_id),
	).Update(FavoriteCnt.Number, FavoriteCnt.Number.Add(-1))
	return err
}

func UpdateCount(video_id int64, cnt int64) error {
	_, err := DoFavoriteCnt.Where(
		FavoriteCnt.VideoId.Eq(video_id),
	).Update(FavoriteCnt.Number, cnt)
	return err
}

func UpdateCacheCount(video_id int64, is_fav bool) error {
	var op int
	if is_fav {
		op = 1
	} else {
		op = -1
	}

	curr_str, err := RedisClients[misc.FavCntCache].Get(context.Background(), fmt.Sprint(video_id))

	if err == redis.Nil {
		return RedisClients[misc.FavCntCache].Set(context.Background(), fmt.Sprint(video_id), "1", -1)
	}

	curr, _ := strconv.Atoi(curr_str)
	curr += op

	return RedisClients[misc.FavCntCache].Set(context.Background(), fmt.Sprint(video_id), fmt.Sprint(curr), -1)
}
