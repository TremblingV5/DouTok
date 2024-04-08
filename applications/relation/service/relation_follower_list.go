package service

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/dal/model"
	"github.com/TremblingV5/DouTok/applications/relation/dal/query"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RelationFollowerListService struct {
	ctx context.Context
}

func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{ctx: ctx}
}

func (s *RelationFollowerListService) RelationFollowerList(req *relation.DouyinRelationFollowerListRequest) (error, []*user.User) {
	// 从 cache 读
	err, follower := ReadFollowerListFromCache(fmt.Sprintf("%d", req.UserId))
	if err != nil || follower == nil {
		klog.Errorf("read follower list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowerListFromDB(req.UserId)
		if err != nil {
			klog.Errorf("read follower list from db error, err = %s", err)
			return err, nil
		} else {
			// 添加 cache
			err := WriteFollowerListToCache(fmt.Sprintf("%d", req.UserId), relationList)
			if err != nil {
				klog.Errorf("update follower list to cache error, err = %s", err)
			}
			// 为 follower 赋值
			list := make([]int64, len(relationList))
			for _, v := range relationList {
				list = append(list, v.UserId)
			}
			follower = list
		}
	}
	// 去用户服务查询 follow list 的 user 信息
	request := new(user.DouyinUserListRequest)
	request.UserList = follower
	resp, err := rpc.GetUserListByIds(context.Background(), request)
	if err != nil {
		return err, nil
	}
	return nil, resp.GetUserList()
}

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
