package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
)

func PackageRelationFriendListResponse(ctx context.Context, result *relationDomain.DoutokListRelationResponse, userId int64, err error) (*relation.DouyinRelationFriendListResponse, error) {
	if err != nil {
		return nil, err
	}

	var friendUserList []*entity.FriendUser
	for _, v := range result.UserList {
		messageInfo, err := rpc.MessageDomainRPCClient.ListMessage(ctx, &messageDomain.DoutokListMessageRequest{
			UserId:   userId,
			ToUserId: v.GetId(),
		})
		if err != nil {
			continue
		}
		var msgType int64
		if len(messageInfo.Message) > 0 {
			if messageInfo.Message[len(messageInfo.Message)-1].ToUserId == userId {
				msgType = 0
			} else {
				msgType = 1
			}
			friendUserList = append(friendUserList, &entity.FriendUser{
				User:    v,
				Message: messageInfo.Message[len(messageInfo.Message)-1].Content,
				MsgType: msgType,
			})
		}
		friendUserList = append(friendUserList, &entity.FriendUser{
			User:    v,
			Message: "",
			MsgType: 0,
		})
	}

	return &relation.DouyinRelationFriendListResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserList:   friendUserList,
	}, err
}
