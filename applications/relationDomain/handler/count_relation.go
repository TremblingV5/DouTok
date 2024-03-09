package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *RelationDomainServiceImpl) CountRelation(ctx context.Context, req *relationDomain.DoutokCountRelationRequest) (resp *relationDomain.DoutokCountRelationResponse, err error) {
	resp = &relationDomain.DoutokCountRelationResponse{
		Result: make(map[int64]int64),
	}

	/*
		0 -> 关注数 1 -> 粉丝数
	*/
	if req.ActionType == 0 {
		for _, v := range req.UserId {
			follow, err := service.NewRelationCountService(ctx).GetFollowCount(v)
			if err != nil {
				continue
			}
			resp.Result[v] = follow
		}
	} else if req.ActionType == 1 {
		for _, v := range req.UserId {
			follower, err := service.NewRelationCountService(ctx).GetFollowerCount(v)
			if err != nil {
				continue
			}
			resp.Result[v] = follower
		}
	}

	pack.BuildRelationCountResp(errno.Success, resp)
	return resp, nil
}
