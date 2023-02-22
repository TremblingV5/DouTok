package service

import (
	"context"
	"fmt"

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
	err, friendList := GetFriendList(req.UserId)
	if err != nil {
		return err, nil
	}
	// 去用户服务查询 friendList 的 user 信息
	reqUser := new(user.DouyinUserListRequest)
	reqUser.UserList = friendList
	respUser, err := rpc.GetUserListByIds(context.Background(), reqUser)
	if err != nil {
		return err, nil
	}
	// 去 message 服务查询对应好友列表的最新消息 返回一个 map
	reqMsg := new(message.DouyinFriendListMessageRequest)
	reqMsg.UserId = req.UserId
	reqMsg.FriendIdList = friendList
	respMsg, err := rpc.GetFriendList(context.Background(), reqMsg)

	for k, v := range respMsg.Result {
		klog.Infof("res key = %d, msg = %s\n", k, v.Content)
	}

	if err != nil {
		return err, nil
	}
	fList := make([]*relation.FriendUser, 0, len(reqUser.GetUserList()))
	for _, v := range respUser.GetUserList() {
		// 0为当前请求用户接受的消息，1为当前请求用户发送的消息
		msgType := 0
		if respMsg.Result[v.Id].FromUserId == req.UserId {
			msgType = 1
		}

		klog.Infof("user_id = %s, msgType = %d\n", respMsg.Result[v.Id].Content, int64(msgType))

		friend := &relation.FriendUser{
			Id:            v.Id,
			Name:          v.Name,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      v.IsFollow,
			Avatar:        v.Avatar,
			Message:       respMsg.Result[v.Id].Content,
			MsgType:       int64(msgType),
		}
		fList = append(fList, friend)
	}
	return nil, fList
}

// 查数据库
func GetFriendList(user_id int64) (error, []int64) {
	followMap := make(map[int64]bool)
	// 获取 follow
	err, follow := ReadFollowListFromCache(fmt.Sprintf("%d", user_id))
	if err != nil || follow == nil {
		klog.Errorf("read follow list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowListFromDB(user_id)
		if err != nil {
			klog.Errorf("read follow list from db error, err = %s", err)
			return err, nil
		} else {
			// 添加 cache
			err := WriteFollowListToCache(fmt.Sprintf("%d", user_id), relationList)
			if err != nil {
				klog.Errorf("update follow list to cache error, err = %s", err)
			}
			// 为 follow 赋值
			list := make([]int64, len(relationList))
			for _, v := range relationList {
				list = append(list, v.ToUserId)
			}
			follow = list
		}
	}
	// 为 map 赋值
	for _, v := range follow {
		followMap[v] = true
	}
	// 获取 follower
	err, follower := ReadFollowerListFromCache(fmt.Sprintf("%d", user_id))
	if err != nil || follower == nil {
		klog.Errorf("read follower list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowerListFromDB(user_id)
		if err != nil {
			klog.Errorf("read follower list from db error, err = %s", err)
			return err, nil
		} else {
			// 添加 cache
			err := WriteFollowerListToCache(fmt.Sprintf("%d", user_id), relationList)
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
	friendList := make([]int64, 0)
	for _, v := range follower {
		if followMap[v] == true {
			friendList = append(friendList, v)
		}
	}
	return nil, friendList
}
