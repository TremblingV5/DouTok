package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/messageDomain/pack"
	"github.com/TremblingV5/DouTok/applications/messageDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func (s *MessageDomainServiceImpl) ListMessage(ctx context.Context, req *messageDomain.DoutokListMessageRequest) (resp *messageDomain.DoutokListMessageResponse, err error) {
	// 从 hbase 返回历史消息列表（会话id的概念）
	resp = new(messageDomain.DoutokListMessageResponse)

	err, messageList := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		pack.BuildMessageChatResp(err, resp)
		return resp, nil
	}
	resp.Message = messageList

	pack.BuildMessageChatResp(errno.Success, resp)
	return resp, nil
}
