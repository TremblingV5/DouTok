package service

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MessageFriendService struct {
	ctx context.Context
}

func NewMessageFriendService(ctx context.Context) *MessageFriendService {
	return &MessageFriendService{ctx: ctx}
}

func (s *MessageFriendService) MessageFriendList(req *message.DouyinFriendListMessageRequest) (error, map[int64]*message.Message) {
	// 从 redis 缓存读 key:会话id value: 消息内容
	result := make(map[int64]*message.Message)
	for _, friendId := range req.GetFriendIdList() {
		sessionId := utils.GenerateSessionId(req.UserId, friendId)
		content, err := RedisClient.HGet(context.Background(), sessionId, "content").Result()
		fromUserId, err := RedisClient.HGet(context.Background(), sessionId, "from_user_id").Float64()
		toUserId, err := RedisClient.HGet(context.Background(), sessionId, "to_user_id").Float64()
		if err != nil {
			klog.Errorf("get friend list message error, sessionId = %s, err = %s", sessionId, err)
			// TODO 从 hbase获取最新一条聊天记录（慢）应该使用基于 redis 的存储（集群确保可用性）
		}

		message := message.Message{
			Content:    content,
			FromUserId: int64(fromUserId),
			ToUserId:   int64(toUserId),
		}
		result[friendId] = &message
	}
	return nil, result
}
