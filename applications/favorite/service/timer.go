package service

import (
	"strconv"
	"time"

	"github.com/TremblingV5/DouTok/pkg/dlog"
)

func UpdateFavMap() {
	time.Sleep(time.Duration(5) * time.Second)

	keyList := []string{}

	FavCount.Iter(func(key string, v interface{}) {
		keyList = append(keyList, key)

		key_i64, _ := strconv.ParseInt(key, 10, 64)
		err := UpdateCount(key_i64, v.(int64))
		if err != nil {
			// TODO: 写日志
			dlog.Warn("Write favourite count to RDB defeat: " + key + " with count: " + v.(string))
		}

		FavCount.Set(key, 0)

		err = DelCount2Cache(key_i64)
		if err != nil {
			dlog.Warn("Delete favourite count from third party cache defeat: " + key)
		}
	})

	for _, v := range keyList {
		FavCount.Set(v, 0)
	}
}
