package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/hbaseHandle"
	"github.com/TremblingV5/DouTok/pkg/misc"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *message.DouyinMessageChatRequest) (error, []*message.Message) {
	// 从 hbase 获取聊天记录
	messageList := make([]*message.Message, 0)
	sessionId := utils.GenerateSessionId(req.UserId, req.ToUserId)
	res, err := HBClient.Scan(ViperConfig.Viper.GetString("Hbase.Table"), hbaseHandle.GetFilterByRowKeyPrefix(sessionId)...)
	if err != nil {
		return err, nil
	}
	for _, v := range res {
		packMsg := pack.Message{}
		err := misc.Map2Struct4HB(v, &packMsg)
		if err != nil {
			continue
		}
		message := message.Message{
			Id:         packMsg.Id,
			ToUserId:   packMsg.ToUserId,
			FromUserId: packMsg.FromUserId,
			Content:    packMsg.Content,
			ActionType: packMsg.ContentType,
			CreateTime: packMsg.CreateTime,
		}
		messageList = append(messageList, &message)
	}
	return nil, messageList
}
