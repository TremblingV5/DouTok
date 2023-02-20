package service

import (
	"bou.ke/monkey"
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestRelationFollowListService(t *testing.T) {
	Init()
	relService := NewRelationFollowListService(context.Background())
	req := &relation.DouyinRelationFollowListRequest{
		UserId: 10001000,
	}

	monkey.Patch(rpc.GetUserListByIds, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*user.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
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
	req = &relation.DouyinRelationFollowListRequest{
		UserId: 10006000,
	}

	monkey.Patch(rpc.GetUserListByIds, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*user.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
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
