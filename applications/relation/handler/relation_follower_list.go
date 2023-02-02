package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	k := utils.KeyGen(req.UserId, 2, 2)
	result, err := redis.RD.SGetObj(ctx, k)
	if err != nil {

		return nil, err
	}
	//缓存命中
	if len(result) != 0 {
		return &relation.DouyinRelationFollowerListResponse{UserList: result, StatusCode: 0, StatusMsg: "success"}, nil
	}
	//查数据库
	result, err = db.GetFollowerList(req.UserId)
	if err != nil {
		return
	}
	//写入缓存
	redis.RD.SAddObj(ctx, k, result...)
	return &relation.DouyinRelationFollowerListResponse{UserList: result, StatusMsg: "success", StatusCode: 0}, nil

}
