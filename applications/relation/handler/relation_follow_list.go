package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	k := utils.KeyGen(req.UserId, 1, 2)
	result, err := redis.RD.SGetObj(ctx, k)
	if err != nil {
		return nil, err
	}
	//命中缓存
	if len(result) != 0 {
		return &relation.DouyinRelationFollowListResponse{StatusCode: 0, StatusMsg: "success", UserList: result}, nil
	}
	//未命中缓存,查数据库，跟新缓存
	users, err := db.GetFollowList(req.UserId)
	if err != nil {
		return nil, err
	}
	//跟新缓存
	go redis.RD.SAddObj(ctx, k, users...)

	return &relation.DouyinRelationFollowListResponse{StatusMsg: "success", StatusCode: 0, UserList: users}, nil
}
