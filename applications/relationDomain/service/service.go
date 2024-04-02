package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (s *Service) AddRelation(ctx context.Context, userId, toUserId int64) error {

	followList, err := s.followListRedis.Get(ctx, userId)
	if err != nil {
		return err
	}

	isFollowed := false

	for _, v := range followList {
		if v == toUserId {
			isFollowed = true
			break
		}
	}

	if !isFollowed {
		isFollowed, err = ReadIsFollowFromDB(userId, toUserId)
		if err != nil {
			isFollowed = false
		}
	}

	if isFollowed {
		return nil
	}

	// 在 SafeMap 中更新局部关注数和粉丝数
	followKey := fmt.Sprintf("%s%d", constants.FollowCount, userId)
	followerKey := fmt.Sprintf("%s%d", constants.FollowerCount, toUserId)
	follow, ok := ConcurrentMap.Get(followKey)
	if !ok {
		klog.Infof("get follow count from concurrentMap false")
	}
	follower, ok := ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Infof("get follow count from concurrentMap false")
	}
	op := int64(1)
	// TODO 如果关注或者取关对应的增加 safemap 值，前提是需要验证重复性操作
	mu.Lock()
	if follow == nil {
		klog.Infof("set follow %s, %d\n", followKey, op)
		ConcurrentMap.Set(followKey, op)
	} else {
		klog.Infof("set follow %s, %d\n", followKey, follow.(int64)+op)
		ConcurrentMap.Set(followKey, follow.(int64)+op)
	}
	if follower == nil {
		klog.Infof("set follower %s, %d\n", followerKey, op)
		ConcurrentMap.Set(followerKey, op)
	} else {
		klog.Infof("set follower %s, %d\n", followerKey, follower.(int64)+op)
		ConcurrentMap.Set(followerKey, follower.(int64)+op)
	}
	follow, ok = ConcurrentMap.Get(followKey)
	if !ok {
		klog.Errorf("concurrentMap get false")
	}
	follower, ok = ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Errorf("concurrentMap get false")
	}
	klog.Infof("%s follow = %d\n", followKey, follow.(int64))
	klog.Infof("%s follower = %d\n", followerKey, follower.(int64))
	mu.Unlock()
	// 使用同步producer，将消息存入 kafka
	// 构建消息
	val, err := json.Marshal(NewRelation(userId, toUserId, 0))
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: ViperConfig.Viper.GetStringSlice("Kafka.Topics")[0],
		Value: sarama.StringEncoder(val),
	}
	partition, offset, err := SyncProducer.SendMessage(msg)

	if err == nil {
		klog.Infof("produce success, partition: %d, offset: %d\n", partition, offset)
	} else {
		return err
	}

	return nil
}

func (s *Service) RmRelation(ctx context.Context, userId, toUserId int64) error {
	// 从 SafeMap 中更新局部关注数和粉丝数
	followKey := fmt.Sprintf("%s%d", constants.FollowCount, userId)
	followerKey := fmt.Sprintf("%s%d", constants.FollowerCount, toUserId)
	follow, ok := ConcurrentMap.Get(followKey)
	if !ok {
		klog.Infof("get follow count from concurrentMap false")
	}
	follower, ok := ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Infof("get follow count from concurrentMap false")
	}
	op := int64(-1)
	// TODO 如果关注或者取关对应的增加 safemap 值，前提是需要验证重复性操作
	mu.Lock()
	if follow == nil {
		klog.Infof("set follow %s, %d\n", followKey, op)
		ConcurrentMap.Set(followKey, op)
	} else {
		klog.Infof("set follow %s, %d\n", followKey, follow.(int64)+op)
		ConcurrentMap.Set(followKey, follow.(int64)+op)
	}
	if follower == nil {
		klog.Infof("set follower %s, %d\n", followerKey, op)
		ConcurrentMap.Set(followerKey, op)
	} else {
		klog.Infof("set follower %s, %d\n", followerKey, follow.(int64)+op)
		ConcurrentMap.Set(followerKey, follower.(int64)+op)
	}
	follow, ok = ConcurrentMap.Get(followKey)
	if !ok {
		klog.Errorf("concurrentMap get false")
	}
	follower, ok = ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Errorf("concurrentMap get false")
	}
	klog.Infof("%s follow = %d\n", followKey, follow.(int64))
	klog.Infof("%s follower = %d\n", followerKey, follower.(int64))
	mu.Unlock()
	// 使用同步producer，将消息存入 kafka
	// 构建消息
	val, err := json.Marshal(NewRelation(userId, toUserId, 1))
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: ViperConfig.Viper.GetStringSlice("Kafka.Topics")[0],
		Value: sarama.StringEncoder(val),
	}
	partition, offset, err := SyncProducer.SendMessage(msg)

	if err == nil {
		klog.Infof("produce success, partition: %d, offset: %d\n", partition, offset)
	} else {
		return err
	}

	return nil
}

func (s *Service) ListFollowList(ctx context.Context, userId int64) ([]*entity.User, error) {
	// 从 cache 读
	follow, err := s.followListRedis.Get(ctx, userId)
	if err != nil || follow == nil {
		klog.Errorf("read follow list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowListFromDB(userId)
		if err != nil {
			klog.Errorf("read follow list from db error, err = %s", err)
			return nil, err
		} else {
			// 添加 cache
			err := WriteFollowListToCache(fmt.Sprintf("%d", userId), relationList)
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

	// 去用户服务查询 follow list 的 user 信息
	// request := new(userDomain.DoutokGetUserInfoRequest)
	// request.UserId = follow
	// resp, err := rpc.UserDomainRPCClient.GetUserInfo(context.Background(), request)
	// if err != nil {
	// 	return nil, err
	// }

	// var result []*entity.User
	// for _, v := range resp.UserList {
	// 	result = append(result, v)
	// }
	result := make([]*entity.User, 1)
	for _, v := range follow {
		result = append(result, &entity.User{
			Id: v,
		})
	}

	return result, nil
}

func (s *Service) ListFollowerList(ctx context.Context, userId int64) ([]*entity.User, error) {
	// 从 cache 读
	follower, err := s.followerListRedis.Get(ctx, userId)
	if err != nil || follower == nil {
		klog.Errorf("read follower list from cache error, err = %s", err)
		// 从 db 读
		err, relationList := ReadFollowerListFromDB(userId)
		if err != nil {
			klog.Errorf("read follower list from db error, err = %s", err)
			return nil, err
		} else {
			// 添加 cache
			err := s.followerListRedis.Set(ctx, userId, relationList)
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
	// request := new(userDomain.DoutokGetUserInfoRequest)
	// request.UserId = follower
	// resp, err := rpc.UserDomainRPCClient.GetUserInfo(context.Background(), request)
	// if err != nil {
	// 	return nil, err
	// }

	// var result []*entity.User
	// for _, v := range resp.UserList {
	// 	result = append(result, v)
	// }

	result := make([]*entity.User, 0)
	for _, v := range follower {
		result = append(result, &entity.User{
			Id: v,
		})
	}

	return result, nil
}

func (s *Service) ListFriendList(ctx context.Context, userId int64) ([]*entity.User, error) {
	// 从 cache 读
	err, friendList := GetFriendList(userId)
	if err != nil {
		return nil, err
	}
	// 去用户服务查询 friendList 的 user 信息
	// reqUser := new(userDomain.DoutokGetUserInfoRequest)
	// reqUser.UserId = friendList
	// respUser, err := rpc.UserDomainRPCClient.GetUserInfo(context.Background(), reqUser)
	// if err != nil {
	// 	return nil, err
	// }
	// 去 message 服务查询对应好友列表的最新消息 返回一个 map
	reqMsg := new(relationDomain.DoutokListRelationRequest)
	reqMsg.UserId = userId
	reqMsg.ActionType = 2
	// reqMsg.FriendIdList = friendList
	// _, err = rpc.RelationDomainRPCClient.ListRelation(context.Background(), reqMsg)

	// for k, v := range respMsg.UserList {
	// 	klog.Infof("res key = %d, msg = %s\n", k, v.Content)
	// }

	if err != nil {
		return nil, err
	}
	var fList []*entity.User
	for _, v := range friendList {
		// 0为当前请求用户接受的消息，1为当前请求用户发送的消息
		// msgType := 0
		// if respMsg.UserList[v.Id].FromUserId == userId {
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
	return fList, nil
}

func (s *Service) GetFollowCount(ctx context.Context, userId int64) (int64, error) {
	follow, err := s.followCountRedis.Get(ctx, userId)
	if err != nil || follow == 0 {
		// 记录日志
		klog.Errorf("read follow count from cache error, err = %s", err)
		// 读 db 获取关注数
		err, follow = ReadFollowCountFromDB(userId)
		if err != nil {
			// 记录日志
			klog.Errorf("read follow count from db error, err = %s", err)
			follow = 0
		}
		// 新增 cache 关注数
		err = s.followCountRedis.Set(ctx, userId, follow)
		if err != nil {
			// 记录日志
			klog.Errorf("update follow count to cache error, err = %s", err)
		}
	}
	return follow, nil
}

func (s *Service) GetFollowerCount(ctx context.Context, userId int64) (int64, error) {
	follower, err := s.followerCountRedis.Get(ctx, userId)
	if err != nil || follower == 0 {
		// 记录日志
		klog.Errorf("read follower count from cache error, err = %s", err)
		// 读 db 获取粉丝数
		err, follower = ReadFollowerCountFromDB(userId)
		if err != nil {
			// 记录日志
			klog.Errorf("read follower count from db error, err = %s", err)
			follower = 0
		}
		// 新增 cache 粉丝数
		err = s.followerCountRedis.Set(ctx, userId, follower)
		if err != nil {
			// 记录日志
			klog.Errorf("update follower count to cache error, err = %s", err)
		}
	}
	return follower, nil
}

func (s *Service) GetFriendCount(ctx context.Context, userId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}
