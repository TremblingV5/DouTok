package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
)

func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinRelationActionRequest) (resp *message.DouyinRelationActionResponse, err error) {
	if ok, msg := check(req); ok {
		resp, _ := service.PackMessageActionRes(1, msg)
		return resp, nil
	}

	if err := service.SaveMessage(req.UserId, req.ToUserId, req.Content); err != nil {
		resp, _ := service.PackMessageActionRes(1, err.Error())
		return resp, err
	}

	resp, _ = service.PackMessageActionRes(0, err.Error())
	return resp, nil
}

func check(req *message.DouyinRelationActionRequest) (bool, string) {
	if len(req.Content) == 0 {
		return true, "缺少聊天内容"
	}

	if req.ActionType != 1 {
		return true, "指令错误"
	}

	return false, ""
}
