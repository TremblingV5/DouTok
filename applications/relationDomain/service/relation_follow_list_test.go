package service

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestRelationFollowListService(t *testing.T) {
	Init()
	relService := NewRelationFollowListService(context.Background())
	req := &relationDomain.DoutokListRelationRequest{
		UserId: 10001000,
	}

	monkey.Patch(rpc.UserDomainRPCClient.GetUserInfo, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*entity.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		return &user.DouyinUserListResponse{
			UserList: userList,
		}, nil
	})

	err, ret := relService.RelationFollowList(req)
	assert.Nil(t, err)
	for _, user := range ret {
		println(user.Id, user.Name)
	}

	relService = NewRelationFollowListService(context.Background())
	req = &relationDomain.DoutokListRelationRequest{
		UserId: 10006000,
	}

	monkey.Patch(rpc.UserDomainRPCClient.GetUserInfo, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*entity.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		return &user.DouyinUserListResponse{
			UserList: userList,
		}, nil
	})

	err, ret = relService.RelationFollowList(req)
	assert.Nil(t, err)
	for _, user := range ret {
		println(user.Id, user.Name)
	}
}
