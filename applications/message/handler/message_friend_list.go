package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/message/dal/model"
	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
)

func (s *MessageServiceImpl) MessageFriendList(ctx context.Context, req *message.DouyinFriendLastMessageRequest) (resp *message.DouyinFriendLastMessageResponse, err error) {
	user_id := req.UserId
	friend_list := req.FriendIdList
	var res map[int64]model.Message
	for _, v := range friend_list {
		list, err := service.QueryMessageFriendListInHBase(user_id, v)
		if err != nil {
			return nil, err
		}
		res[v] = list[0]
	}
	resp, err = service.PackMessageFriendListRes(res, 0, "Success", req)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
