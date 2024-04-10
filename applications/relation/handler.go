package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/TremblingV5/DouTok/applications/relation/service"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationActionResponse)

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		pack.BuildRelationActionResp(err, resp)
		return resp, nil
	}
	pack.BuildRelationActionResp(errno.Success, resp)
	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationFollowListResponse)

	err, followList := service.NewRelationFollowListService(ctx).RelationFollowList(req)
	if err != nil {
		pack.BuildRelationFollowListResp(err, resp)
		return resp, nil
	}
	resp.UserList = followList

	pack.BuildRelationFollowListResp(errno.Success, resp)
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationFollowerListResponse)

	err, followerList := service.NewRelationFollowerListService(ctx).RelationFollowerList(req)
	if err != nil {
		pack.BuildRelationFollowerListResp(err, resp)
		return resp, nil
	}
	resp.UserList = followerList

	pack.BuildRelationFollowerListResp(errno.Success, resp)
	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationFriendListResponse)

	err, friendList := service.NewRelationFriendListService(ctx).RelationFriendList(req)
	if err != nil {
		pack.BuildRelationFriendListResp(err, resp)
		return resp, nil
	}
	resp.UserList = friendList

	pack.BuildRelationFriendListResp(errno.Success, resp)
	return resp, nil
}

// GetFollowCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowCount(ctx context.Context, req *relation.DouyinRelationCountRequest) (resp *relation.DouyinRelationCountResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationCountResponse)

	err, follow, follower := service.NewRelationCountService(ctx).RelationCount(req)
	if err != nil {
		pack.BuildRelationCountResp(err, resp)
		return resp, nil
	}
	resp.FollowCount = follow
	resp.FollowerCount = follower

	pack.BuildRelationCountResp(errno.Success, resp)
	return resp, nil
}
