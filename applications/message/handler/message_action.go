package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/applications/message/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
)

func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (*message.DouyinMessageActionResponse, error) {
	result, err := rpc.MessageDomainClient.AddMessage(ctx, &messageDomain.DoutokAddMessageRequest{
		ToUserId:   req.ToUserId,
		UserId:     req.UserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	return pack.PackageMessageActionResponse(result, err)
}
