package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/query"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/go-redis/redis/v8"
)

// 查缓存
func ReadFollowerListFromCache(user_id string) (error, []int64) {
	res, err := RedisClient.HGetAll(context.Background(), constants.FollowerListPrefix+user_id).Result()
	if err != nil {
		return err, nil
	}
	ret := make([]int64, 0)
	for k, v := range res {
		k_i64, _ := strconv.ParseInt(k, 10, 64)
		if v == "1" {
			ret = append(ret, k_i64)
		}
	}

	if len(ret) <= 0 {
		return redis.Nil, ret
	} else {
		return nil, ret
	}
}

// 查数据库
func ReadFollowerListFromDB(user_id int64) (error, []*model.Relation) {
	res, err := query.Relation.Where(query.Relation.ToUserId.Eq(user_id)).Find()
	if err != nil {
		return err, nil
	}
	return nil, res
}

// 写入缓存
func WriteFollowerListToCache(user_id string, relations []*model.Relation) error {
	val := make([]string, len(relations)*2)
	for _, v := range relations {
		val = append(val, fmt.Sprintf("%d", v.UserId))
		val = append(val, fmt.Sprintf("%d", v.Status))
	}

	_, err := RedisClient.HSet(context.Background(), constants.FollowerListPrefix+user_id, val).Result()
	if err != nil {
		return err
	}
	return nil
}

func WriteFollowerToCache(user_id string, to_user_id string, action_type string) error {
	_, err := RedisClient.HSet(context.Background(), constants.FollowerListPrefix+user_id, to_user_id, action_type).Result()
	return err
}
