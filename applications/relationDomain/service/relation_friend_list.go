package service

import (
	"context"
	"fmt"

	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"

	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationFriendListService struct {
	ctx context.Context
}

func NewRelationFriendListService(ctx context.Context) *RelationFriendListService {
	return &RelationFriendListService{ctx: ctx}
}

func (s *RelationFriendListService) RelationFriendList(req *relationDomain.DoutokListRelationRequest) (error, []*entity.User) {
	// 从 cache 读
	err, friendList := GetFriendList(req.UserId)
	if err != nil {
		return err, nil
	}
	// 去用户服务查询 friendList 的 user 信息
	// reqUser := new(userDomain.DoutokGetUserInfoRequest)
	// reqUser.UserId = friendList
	// respUser, err := rpc.UserDomainRPCClient.GetUserInfo(context.Background(), reqUser)
	// if err != nil {
	// 	return err, nil
	// }
	// 去 message 服务查询对应好友列表的最新消息 返回一个 map
	reqMsg := new(relationDomain.DoutokListRelationRequest)
	reqMsg.UserId = req.UserId
	reqMsg.ActionType = 2
	// reqMsg.FriendIdList = friendList
	// _, err = rpc.RelationDomainRPCClient.ListRelation(context.Background(), reqMsg)

	// for k, v := range respMsg.UserList {
	// 	klog.Infof("res key = %d, msg = %s\n", k, v.Content)
	// }

	if err != nil {
		return err, nil
	}
	var fList []*entity.User
	for _, v := range friendList {
		// 0为当前请求用户接受的消息，1为当前请求用户发送的消息
		// msgType := 0
		// if respMsg.UserList[v.Id].FromUserId == req.UserId {
		// 	msgType = 1
		// }

		// klog.Infof("user_id = %s, msgType = %d\n", respMsg.UserList[v.Id].Content, int64(msgType))

		//friend := &entity.FriendUser{
		//	User: &entity.User{
		//		Id:            v.Id,
		//		Name:          v.Name,
		//		FollowCount:   v.FollowCount,
		//		FollowerCount: v.FollowerCount,
		//		IsFollow:      v.IsFollow,
		//		Avatar:        v.Avatar,
		//	},
		//	Message: respMsg.Result[v.Id].Content,
		//	MsgType: int64(msgType),
		//}
		friend := &entity.User{
			Id: v,
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
