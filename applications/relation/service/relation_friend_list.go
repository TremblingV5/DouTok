package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationFriendListService struct {
	ctx context.Context
}

func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

func (s *RelationFriendListService) RelationFriendList(req *relation.DouyinRelationFriendListRequest) (error, []*relation.FriendUser) {
	// 从 cache 读
	err, follower := ReadFollowerListFromCache(string(req.UserId))
	if err != nil {
		klog.Errorf("read follower list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowerListFromDB(req.UserId)
		if err != nil {
			klog.Errorf("read follower list from db error, err = %s", err)
			return err, nil
		} else {
			// 添加 cache
			err := WriteFollowerListToCache(string(req.UserId), relationList)
			if err != nil {
				klog.Errorf("update follower list to cache error, err = %s", err)
			}
			// 为 follower 赋值
			list := make([]int64, len(relationList))
			for _, v := range relationList {
				list = append(list, v.UserId)
			}
			follower = list
		}
	}
	// 去用户服务查询 follow list 的 user 信息
	reqUser := new(user.DouyinUserListRequest)
	reqUser.UserList = follower
	respUser, err := rpc.GetUserListByIds(context.Background(), reqUser)
	if err != nil {
		return err, nil
	}
	// 去 message 服务查询对应好友列表的最新消息 返回一个 map
	reqMsg := new(message.DouyinFriendListMessageRequest)
	reqMsg.FriendIdList = follower
	respMsg, err := rpc.GetFriendList(context.Background(), reqMsg)
	if err != nil {
		return err, nil
	}
	friendList := make([]*relation.FriendUser, len(reqUser.GetUserList()))
	for _, v := range respUser.GetUserList() {
		user := &user.User{
			Id:            v.Id,
			Name:          v.Name,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      v.IsFollow,
			Avatar:        v.Avatar,
		}
		friend := &relation.FriendUser{
			User:    user,
			Message: respMsg.Result[v.Id].Content,
			MsgType: int64(respMsg.Result[v.Id].ActionType),
		}
		friendList = append(friendList, friend)
	}
	return nil, friendList
}
