package service

import (
	"context"
	"fmt"
	"math"
	"sort"

	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
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
	start := fmt.Sprintf("%s%d", sessionId, req.PreMsgTime)
	end := fmt.Sprintf("%s%d", sessionId, math.MaxInt)
	res, err := HBClient.ScanRange(ViperConfig.Viper.GetString("Hbase.Table"), start, end)
	if err != nil {
		return err, nil
	}
	for _, v := range res {
		hbMsg := pack.HBMessage{}
		err := misc.Map2Struct4HB(v, &hbMsg)
		if err != nil {
			continue
		}
		packMsg := pack.HBMsg2Msg(&hbMsg)
		message := message.Message{
			Id:         packMsg.Id,
			ToUserId:   packMsg.ToUserId,
			FromUserId: packMsg.FromUserId,
			Content:    packMsg.Content,
			CreateTime: packMsg.CreateTime,
		}
		messageList = append(messageList, &message)
	}
	sort.SliceStable(messageList, func(i, j int) bool {
		return messageList[i].CreateTime < messageList[j].CreateTime
	})
	return nil, messageList
}
