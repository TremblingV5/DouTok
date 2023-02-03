package service

import (
	"fmt"

	"github.com/TremblingV5/DouTok/applications/message/dal/model"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
)

func PackMessageActionRes(code int32, msg string) (*message.DouyinRelationActionResponse, error) {
	var resp message.DouyinRelationActionResponse

	resp.StatusCode = code
	resp.StatusMsg = msg

	return &resp, nil
}

func PackMessageChatRes(list []model.Message, code int32, msg string, req *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	res := message.DouyinMessageChatResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	var message_list []*message.Message

	for _, v := range list {
		var temp message.Message

		temp.Content = v.Content
		temp.ToUserId = int64(v.ToUserID)
		temp.CreateTime = fmt.Sprint(v.CreatedAt)
		temp.FromUserId = int64(v.UserID)

		message_list = append(message_list, &temp)
	}

	res.MessageList = message_list

	return &res, nil
}
func PackMessageFriendListRes(list map[int64]model.Message, code int32, msg string, req *message.DouyinFriendLastMessageRequest) (*message.DouyinFriendLastMessageResponse, error) {
	res := message.DouyinFriendLastMessageResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	var message_list map[int64]*message.Message

	for k, v := range list {
		var temp message.Message

		temp.Content = v.Content
		temp.ToUserId = int64(v.ToUserID)
		temp.CreateTime = fmt.Sprint(v.CreatedAt)
		temp.FromUserId = int64(v.UserID)

		message_list[k] = &temp
	}

	res.Result = message_list

	return &res, nil
}
