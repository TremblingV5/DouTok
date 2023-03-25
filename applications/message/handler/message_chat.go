package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/applications/message/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
)

func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	result, err := rpc.MessageDomainClient.ListMessage(ctx, &messageDomain.DoutokListMessageRequest{
		ToUserId:   req.ToUserId,
		UserId:     req.UserId,
		PreMsgTime: req.PreMsgTime,
	})
	return pack.PackageMessageChatResponse(result, err)
}
