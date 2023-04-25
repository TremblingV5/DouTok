package service

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestRelationFollowerListService(t *testing.T) {
	Init()
	relService := NewRelationFollowerListService(context.Background())
	req := &relationDomain.DoutokListRelationRequest{UserId: 10002000, ActionType: 1}

	monkey.Patch(rpc.UserDomainRPCClient.GetUserInfo, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*entity.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		return &user.DouyinUserListResponse{
			UserList: userList,
		}, nil
	})

	err, ret := relService.RelationFollowerList(req)
	assert.Nil(t, err)
	for _, user := range ret {
		println(user.Id, user.Name)
	}

	relService = NewRelationFollowerListService(context.Background())
	req = &relationDomain.DoutokListRelationRequest{UserId: 10007000, ActionType: 1}

	monkey.Patch(rpc.UserDomainRPCClient.GetUserInfo, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *userDomain.DoutokGetUserInfoRequest, err error) {
		// userList := []*entity.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		userList := []int64{10002000, 10003000}
		return &userDomain.DoutokGetUserInfoRequest{
			UserId: userList,
		}, nil
	})

	err, ret = relService.RelationFollowerList(req)
	assert.Nil(t, err)
	for _, user := range ret {
		println(user.Id, user.Name)
	}
}
