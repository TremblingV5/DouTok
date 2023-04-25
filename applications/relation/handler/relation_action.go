package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
)

func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	if req.ActionType == 0 {
		result, err := rpc.RelationDomainRPCClient.AddRelation(ctx, &relationDomain.DoutokAddRelationRequest{
			UserId:   req.UserId,
			ToUserId: req.ToUserId,
		})
		return pack.PackageRelationActionResponse(result.StatusCode, result.StatusMsg, err)
	} else if req.ActionType == 1 {
		result, err := rpc.RelationDomainRPCClient.RmRelation(ctx, &relationDomain.DoutokRmRelationRequest{
			UserId:   req.UserId,
			ToUserId: req.ToUserId,
		})
		return pack.PackageRelationActionResponse(result.StatusCode, result.StatusMsg, err)
	} else {
		return pack.PackageRelationActionResponse(int32(-1), "defeat", err)
	}
}
