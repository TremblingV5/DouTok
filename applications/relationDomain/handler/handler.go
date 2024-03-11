package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (h *RelationDomainHandler) AddRelation(ctx context.Context, req *relationDomain.DoutokAddRelationRequest) (resp *relationDomain.DoutokAddRelationResponse, err error) {
	resp = new(relationDomain.DoutokAddRelationResponse)

	err = h.service.AddRelation(ctx, req.UserId, req.ToUserId)
	if err != nil {
		pack.BuildRelationActionResp(err, resp)
		return resp, nil
	}
	pack.BuildRelationActionResp(errno.Success, resp)
	return resp, nil
}

func (h *RelationDomainHandler) RmRelation(ctx context.Context, req *relationDomain.DoutokRmRelationRequest) (resp *relationDomain.DoutokRmRelationResponse, err error) {
	resp = new(relationDomain.DoutokRmRelationResponse)

	err = h.service.RmRelation(ctx, req.UserId, req.ToUserId)
	if err != nil {
		pack.BuildRmRelationActionResp(err, resp)
		return resp, nil
	}
	pack.BuildRmRelationActionResp(errno.Success, resp)
	return resp, nil
}

func (h *RelationDomainHandler) ListRelation(ctx context.Context, req *relationDomain.DoutokListRelationRequest) (resp *relationDomain.DoutokListRelationResponse, err error) {

	resp = new(relationDomain.DoutokListRelationResponse)
	if req.ActionType == 0 {
		// 关注
		followList, err := h.service.ListFollowList(ctx, req.UserId)
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
		followerList, err := h.service.ListFollowerList(ctx, req.UserId)
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
		friendList, err := h.service.ListFriendList(ctx, req.UserId)
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

func (h *RelationDomainHandler) CountRelation(ctx context.Context, req *relationDomain.DoutokCountRelationRequest) (resp *relationDomain.DoutokCountRelationResponse, err error) {
	resp = &relationDomain.DoutokCountRelationResponse{
		Result: make(map[int64]int64),
	}

	/*
		0 -> 关注数 1 -> 粉丝数
	*/
	if req.ActionType == 0 {
		for _, v := range req.UserId {
			follow, err := h.service.GetFollowCount(ctx, v)
			if err != nil {
				continue
			}
			resp.Result[v] = follow
		}
	} else if req.ActionType == 1 {
		for _, v := range req.UserId {
			follower, err := h.service.GetFollowerCount(ctx, v)
			if err != nil {
				continue
			}
			resp.Result[v] = follower
		}
	}

	pack.BuildRelationCountResp(errno.Success, resp)
	return resp, nil
}
