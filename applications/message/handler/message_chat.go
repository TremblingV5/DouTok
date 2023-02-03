package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
)

func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	user_id := req.UserId
	to_user_id := req.ToUserId

	list, err := service.QueryMessageChatInHBase(user_id, to_user_id)

	if err != nil {
		return nil, err
	}

	resp, err = service.PackMessageChatRes(list, 0, "Success", req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
