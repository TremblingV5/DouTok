package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
)

func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	result, err := rpc.RelationDomainRPCClient.ListRelation(ctx, &relationDomain.DoutokListRelationRequest{
		UserId:     req.UserId,
		ActionType: 2,
	})
	return pack.PackageRelationFriendListResponse(ctx, result, req.UserId, err)
}
