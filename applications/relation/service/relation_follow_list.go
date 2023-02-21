package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/dal/model"
	"github.com/TremblingV5/DouTok/applications/relation/dal/query"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RelationFollowListService struct {
	ctx context.Context
}

func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{ctx: ctx}
}

func (s *RelationFollowListService) RelationFollowList(req *relation.DouyinRelationFollowListRequest) (error, []*user.User) {
	// 从 cache 读
	err, follow := ReadFollowListFromCache(fmt.Sprintf("%d", req.UserId))
	if err != nil || follow == nil {
		klog.Errorf("read follow list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowListFromDB(req.UserId)
		if err != nil {
			klog.Errorf("read follow list from db error, err = %s", err)
			return err, nil
		} else {
			// 添加 cache
			err := WriteFollowListToCache(fmt.Sprintf("%d", req.UserId), relationList)
			if err != nil {
				klog.Errorf("update follow list to cache error, err = %s", err)
			}
			// 为 follow 赋值
			list := make([]int64, len(relationList))
			for _, v := range relationList {
				list = append(list, v.ToUserId)
			}
			follow = list
		}
	}
	// 去用户服务查询 follow list 的 user 信息
	request := new(user.DouyinUserListRequest)
	request.UserList = follow
	resp, err := rpc.GetUserListByIds(context.Background(), request)
	if err != nil {
		return err, nil
	}
	return nil, resp.GetUserList()
}

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
