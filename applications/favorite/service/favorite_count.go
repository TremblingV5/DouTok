package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
	"strconv"
	"time"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/go-redis/redis/v8"
)

func QueryFavoriteCount(video_id []int64) (map[int64]int64, error) {
	resMap := make(map[int64]int64)

	// 1. 从内存中查找喜欢数
	find_again := []int64{}

	for _, v := range video_id {
		cnt, ok, err := ReadFavTotalCount(v)
		if err != nil {
			return nil, err
		}

		if !ok {
			find_again = append(find_again, v)
		} else {
			resMap[v] = cnt
		}
	}

	find_again_again := []int64{}

	// 2. 从Redis中查找喜欢数
	for _, v := range find_again {
		cnt, ok, err := ReadCountFromCache(v)
		if err != nil {
			return nil, err
		}

		if !ok {
			find_again_again = append(find_again_again, v)
		} else {
			resMap[v] = cnt
			FavTotalCount.Set(fmt.Sprint(v), cnt)
		}
	}

	// 3. 从MySQL中查找喜欢数
	res, err := DoFavoriteCnt.Where(
		FavoriteCnt.VideoId.In(find_again_again...),
	).Find()

	if err != nil {
		return nil, err
	}

	for _, v := range res {
		resMap[v.VideoId] = v.Number
		WriteCount2Cache(v.VideoId, v.Number)
	}

	// 4. 如果仍然没有查找到该记录，则置0
	for _, v := range video_id {
		if _, ok := resMap[v]; !ok {
			resMap[v] = 0
		}
	}

	return resMap, nil
}

func ReadCountFromCache(video_id int64) (int64, bool, error) {
	data, err := RedisClients[misc.FavTotalCntCache].Get(context.Background(), fmt.Sprint(video_id))

	if err == redis.Nil {
		return -1, false, nil
	} else if err != nil {
		return -1, false, err
	}

	data_i, _ := strconv.Atoi(data)

	return int64(data_i), true, nil
}

func ReadFavTotalCount(videoId int64) (int64, bool, error) {
	data, ok := FavTotalCount.Get(fmt.Sprint(videoId))
	if !ok {
		return 0, false, nil
	}

	return int64(data.(int)), true, nil
}

func WriteCount2Cache(video_id int64, cnt int64) error {
	return RedisClients[misc.FavCntCache].Set(context.Background(), fmt.Sprint(video_id), fmt.Sprint(cnt), time.Second*60*60*1)
}

func DelCount2Cache(video_id int64) error {
	return RedisClients[misc.FavCntCache].Del(context.Background(), fmt.Sprint(video_id))
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
	data, _ := DoFavoriteCnt.Where(
		FavoriteCnt.VideoId.Eq(video_id),
	).First()

	if data != nil {
		_, err := DoFavoriteCnt.Where(
			FavoriteCnt.VideoId.Eq(video_id),
		).Update(FavoriteCnt.Number, FavoriteCnt.Number.Add(cnt))

		if err != nil {
			return err
		}
	} else {
		err := DoFavoriteCnt.Create(
			&model.FavoriteCount{
				VideoId: video_id,
				Number:  cnt,
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
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
