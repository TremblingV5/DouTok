package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"

	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

// 查缓存
func ReadFollowListFromCache(user_id string) (error, []int64) {
	res, err := RedisClient.HGetAll(context.Background(), constants.FollowListPrefix+user_id).Result()
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
func ReadFollowListFromDB(user_id int64) (error, []*model.Relation) {
	res, err := query.Relation.Where(query.Relation.UserId.Eq(user_id)).Find()
	if err != nil {
		return err, nil
	}
	return nil, res
}

func ReadIsFollowFromDB(userId int64, toUserId int64) (bool, error) {
	res, err := query.Relation.Where(
		query.Relation.UserId.Eq(userId),
		query.Relation.ToUserId.Eq(toUserId),
	).First()

	if err != nil {
		return false, err
	}

	if res.Status == 1 {
		return false, nil
	} else {
		return true, nil
	}
}

// 写入缓存
func WriteFollowListToCache(user_id string, relations []*model.Relation) error {
	val := make([]string, len(relations)*2)
	for _, v := range relations {
		val = append(val, fmt.Sprintf("%d", v.ToUserId))
		val = append(val, fmt.Sprintf("%d", v.Status))
	}

	_, err := RedisClient.HSet(context.Background(), constants.FollowListPrefix+user_id, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// 写入 DB
func WriteFollowToDB(rel *pack.Relation) error {
	res, err := query.Relation.Where(
		query.Relation.UserId.Eq(rel.UserId),
		query.Relation.ToUserId.Eq(rel.ToUserId),
	).Find()
	if err != nil {
		return err
	}
	if len(res) > 0 {
		// 已经存在关注关系
		_, err := query.Relation.Where(
			query.Relation.UserId.Eq(rel.UserId),
			query.Relation.ToUserId.Eq(rel.ToUserId),
		).Update(
			query.Relation.Status, rel.ActionType,
		)
		if err != nil {
			return err
		}
	} else {
		// 不存在则插入
		id := utils.GetSnowFlakeId().Int64()
		err := query.Relation.Create(
			&model.Relation{
				ID:       id,
				UserId:   rel.UserId,
				ToUserId: rel.ToUserId,
				Status:   int(rel.ActionType),
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// 单条写入/更新缓存
func WriteFollowToCache(user_id string, to_user_id string, action_type string) error {
	_, err := RedisClient.HSet(context.Background(), constants.FollowListPrefix+user_id, to_user_id, action_type).Result()
	return err
}
