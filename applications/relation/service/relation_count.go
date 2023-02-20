package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/dal/model"
	"github.com/TremblingV5/DouTok/applications/relation/dal/query"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationCountService struct {
	ctx context.Context
}

func NewRelationCountService(ctx context.Context) *RelationCountService {
	return &RelationCountService{ctx: ctx}
}

func (s *RelationCountService) RelationCount(req *relation.DouyinRelationCountRequest) (error, int64, int64) {

	// 读 cache 获取关注数
	err, follow := ReadFollowCountFromCache(fmt.Sprintf("%d", req.UserId))
	if err != nil || follow == 0 {
		// 记录日志
		klog.Errorf("read follow count from cache error, err = %s", err)
		// 读 db 获取关注数
		err, follow = ReadFollowCountFromDB(req.UserId)
		if err != nil {
			// 记录日志
			klog.Errorf("read follow count from db error, err = %s", err)
			follow = 0
		}
		// 新增 cache 关注数
		err = WriteFollowCountToCache(fmt.Sprintf("%d", req.UserId), follow)
		if err != nil {
			// 记录日志
			klog.Errorf("update follow count to cache error, err = %s", err)
		}
	}
	// 读 cache 获取粉丝数
	err, follower := ReadFollowerCountFromCache(fmt.Sprintf("%d", req.UserId))
	if err != nil || follower == 0 {
		// 记录日志
		klog.Errorf("read follower count from cache error, err = %s", err)
		// 读 db 获取粉丝数
		err, follower = ReadFollowerCountFromDB(req.UserId)
		if err != nil {
			// 记录日志
			klog.Errorf("read follower count from db error, err = %s", err)
			follower = 0
		}
		// 新增 cache 粉丝数
		err = WriteFollowerCountToCache(fmt.Sprintf("%d", req.UserId), follower)
		if err != nil {
			// 记录日志
			klog.Errorf("update follower count to cache error, err = %s", err)
		}
	}
	return nil, follow, follower
}

func ReadFollowCountFromDB(user_id int64) (error, int64) {
	res, err := query.FollowCount.Where(query.FollowCount.UserId.Eq(user_id)).First()
	if err != nil {
		return err, 0
	}
	return nil, res.Number
}

func ReadFollowCountFromCache(user_id string) (error, int64) {
	ret := RedisClient.HGet(context.Background(), user_id, constants.FollowCount)
	err := ret.Err()
	if err != nil {
		return err, 0
	}
	follow, err := ret.Int64()
	if err != nil {
		return err, 0
	}
	return nil, follow
}

func WriteFollowCountToCache(user_id string, follow int64) error {
	ret := RedisClient.HSet(context.Background(), user_id, map[string]interface{}{constants.FollowCount: follow})
	err := ret.Err()
	if err != nil {
		return err
	}
	return nil
}

func UpdateFollowCountFromDB(user_id int64, op int64) error {
	res, err := query.FollowCount.Where(
		query.FollowCount.UserId.Eq(user_id),
	).Find()
	if err != nil {
		return err
	}
	if len(res) > 0 {
		// 已经存在
		_, err := query.FollowCount.Where(
			query.FollowCount.UserId.Eq(user_id),
		).Update(query.FollowCount.Number, query.FollowCount.Number.Add(op))
		return err
	} else {
		err := query.FollowCount.Create(
			&model.FollowCount{
				UserId: user_id,
				Number: op,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateFollowerCountFromDB(user_id int64, op int64) error {
	res, err := query.FollowerCount.Where(
		query.FollowerCount.UserId.Eq(user_id)).Find()
	if err != nil {
		return err
	}
	if len(res) > 0 {
		_, err := query.FollowerCount.Where(
			query.FollowerCount.UserId.Eq(user_id),
		).Update(query.FollowerCount.Number, query.FollowerCount.Number.Add(op))
		return err
	} else {
		err := query.FollowerCount.Create(
			&model.FollowerCount{
				UserId: user_id,
				Number: op,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteFollowCountCache(user_id string) error {
	_, err := RedisClient.HDel(context.Background(), user_id, constants.FollowCount).Result()
	if err != nil {
		return err
	}
	return nil
}

func ReadFollowerCountFromDB(user_id int64) (error, int64) {
	res, err := query.FollowerCount.Where(query.FollowerCount.UserId.Eq(user_id)).First()
	if err != nil {
		return err, 0
	}
	return nil, res.Number
}

func ReadFollowerCountFromCache(user_id string) (error, int64) {
	ret := RedisClient.HGet(context.Background(), user_id, constants.FollowerCount)
	err := ret.Err()
	if err != nil {
		return err, 0
	}
	follower, err := ret.Int64()
	if err != nil {
		return err, 0
	}
	return nil, follower
}

func WriteFollowerCountToCache(user_id string, follower int64) error {
	ret := RedisClient.HSet(context.Background(), user_id, map[string]interface{}{constants.FollowerCount: follower})
	err := ret.Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteFollowerCountCache(user_id string) error {
	_, err := RedisClient.HDel(context.Background(), user_id, constants.FollowCount).Result()
	if err != nil {
		return err
	}
	return nil
}
