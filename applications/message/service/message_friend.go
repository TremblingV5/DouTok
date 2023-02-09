package service

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/utils"
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
		content := RedisClient.HGet(context.Background(), sessionId, "content").String()
		actionType, _ := RedisClient.HGet(context.Background(), sessionId, "action_type").Int()

		message := message.Message{
			Content:    content,
			ActionType: int32(actionType),
		}
		result[friendId] = &message
	}
	return nil, result
}
