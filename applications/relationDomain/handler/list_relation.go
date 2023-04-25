package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/applications/relationDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *RelationDomainServiceImpl) ListRelation(ctx context.Context, req *relationDomain.DoutokListRelationRequest) (resp *relationDomain.DoutokListRelationResponse, err error) {
	resp = new(relationDomain.DoutokListRelationResponse)
	if req.ActionType == 0 {
		// 关注
		err, followList := service.NewRelationFollowListService(ctx).RelationFollowList(req)
		if err != nil {
			pack.BuildRelationFollowListResp(err, resp)
			return resp, nil
		}
		resp.UserList = followList
		pack.BuildRelationFollowListResp(errno.Success, resp)
		return resp, nil
	} else if req.ActionType == 1 {
		// 粉丝
		// 关注
		err, followerList := service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
		if err != nil {
			pack.BuildRelationFollowListResp(err, resp)
			return resp, nil
		}
		resp.UserList = followerList
		pack.BuildRelationFollowListResp(errno.Success, resp)
		return resp, nil
	} else if req.ActionType == 2 {
		// 互相关注
		// 关注
		err, friendList := service.NewRelationFriendListService(ctx).RelationFriendList(req)
		if err != nil {
			pack.BuildRelationFollowListResp(err, resp)
			return resp, nil
		}
		resp.UserList = friendList
		pack.BuildRelationFollowListResp(errno.Success, resp)
		return resp, nil
	}
	return resp, nil
}
