package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/dal/db"
	"github.com/TremblingV5/DouTok/applications/relation/dal/redis"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = &relation.DouyinRelationActionResponse{}

	if req.ActionType == 1 {
		//先跟新数据库
		if err = db.Relation(req.UserId, req.ToUserId); err != nil {
			if err == errno.RelationRepeatedErr {
				resp.StatusMsg = "已关注,不要重复关注"
				resp.StatusCode = 1
			}
			return
		}
		//再删除缓存
		k := redis.Keys(req.UserId, req.ToUserId)
		if err = redis.RD.DelKey(ctx, k...); err != nil {
			return
		}
		resp.StatusMsg = "关注成功"
	} else {
		//先跟新数据库
		if err = db.CancelRelation(req.UserId, req.ToUserId); err != nil {
			if err == errno.RelationUnfollowedErr {
				resp.StatusMsg = "不能取关未关注用户"
				resp.StatusCode = 1
			}
			return
		}
		//再删除缓存
		k := redis.Keys(req.UserId, req.ToUserId)
		if err = redis.RD.DelKey(ctx, k...); err != nil {
			return
		}
		resp.StatusMsg = "取关成功"
	}
	resp.StatusCode = 0
	return
}
