package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/comment/misc"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"strconv"
	"time"
)

var logger = dlog.InitLog(3)

/*
	定时将内存中的局部评论数更新到数据库中，并删除Redis中的评论总数
*/
func UpdateComCountMap() {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		logger.Info("Start iter comment cnt map and update on " + fmt.Sprint(time.Now().Unix()))

		keyList := []string{}

		ComCount.Iter(func(key string, v interface{}) {
			keyList = append(keyList, key)

			keyI64, _ := strconv.ParseInt(key, 10, 64)
			err := UpdateCount(keyI64, int64(v.(int)))
			if err != nil {
				dlog.Warn("Write comment count to RDB defeat: " + key + " with count: " + fmt.Sprint(v.(int)))
			}

			err = DelCount2Cache(key)
			if err != nil {
				dlog.Warn("Delete comment count from third party cache defeat: " + key)
			}
		})

		for _, v := range keyList {
			ComCount.Set(v, 0)
		}
	}
}

/*
	定时将Redis中每个Video的Comment总数更新到内存中的Map
	Redis不存在的视频评论数由单独查询时再添加到Redis中
*/
func UpdateComTotalCntMap() {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		logger.Info("Start iter comment total cnt map and update on " + fmt.Sprint(time.Now().Unix()))

		keyList := []string{}

		ComTotalCount.Iter(func(key string, v interface{}) {
			keyList = append(keyList, key)
		})

		for _, v := range keyList {
			res, err := RedisClients[misc.ComTotalCntCache].Get(context.Background(), fmt.Sprint(v))
			if err != nil {
				continue
			}

			i, _ := strconv.ParseInt(res, 10, 64)
			ComTotalCount.Set(v, i)
		}
	}
}
