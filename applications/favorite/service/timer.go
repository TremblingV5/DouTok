package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"

	"github.com/TremblingV5/DouTok/pkg/dlog"
)

func UpdateFavMap() {
	time.Sleep(time.Duration(5) * time.Second)

	keyList := []string{}

	FavCount.Iter(func(key string, v interface{}) {
		keyList = append(keyList, key)

		key_i64, _ := strconv.ParseInt(key, 10, 64)
		err := UpdateCount(key_i64, int64(v.(int)))
		if err != nil {
			// TODO: 写日志
			dlog.Warn("Write favourite count to RDB defeat: " + key + " with count: " + v.(string))
		}

		err = DelCount2Cache(key_i64)
		if err != nil {
			dlog.Warn("Delete favourite count from third party cache defeat: " + key)
		}
	})

	for _, v := range keyList {
		FavCount.Set(v, 0)
	}
}

func UpdateFavCntMap() {
	time.Sleep(time.Duration(5) * time.Second)

	keyList := []string{}

	FavTotalCount.Iter(func(key string, v interface{}) {
		keyList = append(keyList, key)
	})

	for _, v := range keyList {
		res, err := RedisClients[misc.FavCntCache].Get(context.Background(), fmt.Sprint(v))
		if err != nil {
			continue
		}

		i, _ := strconv.ParseInt(res, 10, 64)
		FavTotalCount.Set(v, i)
	}
}
