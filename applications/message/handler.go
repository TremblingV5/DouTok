package main

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/message/pack"
	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	// 从 hbase 返回历史消息列表（会话id的概念）
	resp = new(message.DouyinMessageChatResponse)

	err, messageList := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		pack.BuildMessageChatResp(err, resp)
		return resp, nil
	}
	resp.MessageList = messageList

	pack.BuildMessageChatResp(errno.Success, resp)
	return resp, nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	resp = new(message.DouyinMessageActionResponse)

	err = service.NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		pack.BuildMessageActionResp(err, resp)
		return resp, nil
	}
	pack.BuildMessageActionResp(errno.Success, resp)
	return resp, nil
}

// MessageFriendList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageFriendList(ctx context.Context, req *message.DouyinFriendListMessageRequest) (resp *message.DouyinFriendListMessageResponse, err error) {
	resp = new(message.DouyinFriendListMessageResponse)

	err, result := service.NewMessageFriendService(ctx).MessageFriendList(req)
	if err != nil {
		pack.BuildMessageFriendResp(err, resp)
		return resp, nil
	}
	resp.Result = result

	pack.BuildMessageFriendResp(errno.Success, resp)
	return resp, nil
}
