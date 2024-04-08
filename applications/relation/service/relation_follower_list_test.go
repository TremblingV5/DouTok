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

func TestRelationFollowerListService(t *testing.T) {
	Init()
	relService := NewRelationFollowerListService(context.Background())
	req := &relation.DouyinRelationFollowerListRequest{UserId: 10002000}

	monkey.Patch(rpc.GetUserListByIds, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*user.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
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
	req = &relation.DouyinRelationFollowerListRequest{UserId: 10007000}

	monkey.Patch(rpc.GetUserListByIds, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*user.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		return &user.DouyinUserListResponse{
			UserList: userList,
		}, nil
	})

	err, ret = relService.RelationFollowerList(req)
	assert.Nil(t, err)
	for _, user := range ret {
		println(user.Id, user.Name)
	}
}
