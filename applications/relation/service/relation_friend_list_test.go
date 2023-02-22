package service

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestRelationFriendListService(t *testing.T) {
	Init()
	relService := NewRelationFriendListService(context.Background())
	req := &relation.DouyinRelationFriendListRequest{UserId: 10001000}

	monkey.Patch(rpc.GetUserListByIds, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*user.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		return &user.DouyinUserListResponse{
			UserList: userList,
		}, nil
	})

	monkey.Patch(rpc.GetFriendList, func(ctx context.Context, req *message.DouyinFriendListMessageRequest) (resp *message.DouyinFriendListMessageResponse, err error) {
		msgMp := make(map[int64]*message.Message)
		msgMp[10002000] = &message.Message{Content: "test msg"}
		msgMp[10003000] = &message.Message{Content: "test msg"}
		return &message.DouyinFriendListMessageResponse{Result: msgMp}, nil
	})

	err, ret := relService.RelationFriendList(req)
	assert.Nil(t, err)
	for _, friend := range ret {
		println(friend.Id, friend.Message)
	}

	relService = NewRelationFriendListService(context.Background())
	req = &relation.DouyinRelationFriendListRequest{UserId: 10008000}

	monkey.Patch(rpc.GetUserListByIds, func(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
		userList := []*user.User{{Id: 10002000, Name: "test2"}, {Id: 10003000, Name: "test3"}}
		return &user.DouyinUserListResponse{
			UserList: userList,
		}, nil
	})

	monkey.Patch(rpc.GetFriendList, func(ctx context.Context, req *message.DouyinFriendListMessageRequest) (resp *message.DouyinFriendListMessageResponse, err error) {
		msgMp := make(map[int64]*message.Message)
		msgMp[10002000] = &message.Message{Content: "test msg"}
		msgMp[10003000] = &message.Message{Content: "test msg"}
		return &message.DouyinFriendListMessageResponse{Result: msgMp}, nil
	})

	err, ret = relService.RelationFriendList(req)
	assert.Nil(t, err)
	for _, friend := range ret {
		println(friend.Id, friend.Message)
	}
}
