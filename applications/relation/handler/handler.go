package handler

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

type Handler struct {
	clients *rpc.Clients
}

func (s *Handler) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	if req.ActionType == 0 {
		result, err := s.clients.Relation.Client.AddRelation(ctx, &relationDomain.DoutokAddRelationRequest{
			UserId:   req.UserId,
			ToUserId: req.ToUserId,
		})
		return &relation.DouyinRelationActionResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	if req.ActionType == 1 {
		result, err := s.clients.Relation.Client.RmRelation(ctx, &relationDomain.DoutokRmRelationRequest{
			UserId:   req.UserId,
			ToUserId: req.ToUserId,
		})
		return &relation.DouyinRelationActionResponse{
			StatusCode: result.StatusCode,
			StatusMsg:  result.StatusMsg,
		}, err
	}

	return &relation.DouyinRelationActionResponse{
		StatusCode: -1,
		StatusMsg:  "unknown action type",
	}, errors.New("unknown action type")
}

func (s *Handler) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	result, err := s.clients.Relation.Client.ListRelation(ctx, &relationDomain.DoutokListRelationRequest{
		UserId:     req.UserId,
		ActionType: 0,
	})

	if err != nil {
		return nil, err
	}

	var userIdList []int64
	for _, v := range result.UserList {
		userIdList = append(userIdList, v.GetId())
	}
	userInfo, err := s.clients.User.Client.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{
		UserId: userIdList,
	})

	var resList []*entity.User
	for _, v := range result.UserList {
		if value, ok := userInfo.UserList[v.GetId()]; ok {
			resList = append(resList, &entity.User{
				Id:              value.GetId(),
				Name:            value.GetName(),
				Avatar:          value.GetAvatar(),
				BackgroundImage: value.GetBackgroundImage(),
				Signature:       value.GetSignature(),
			})
		}
	}

	return &relation.DouyinRelationFollowListResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserList:   resList,
	}, err
}

func (s *Handler) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	result, err := s.clients.Relation.Client.ListRelation(ctx, &relationDomain.DoutokListRelationRequest{
		UserId:     req.UserId,
		ActionType: 1,
	})

	var userIdList []int64
	for _, v := range result.UserList {
		userIdList = append(userIdList, v.GetId())
	}
	userInfo, err := s.clients.User.Client.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{
		UserId: userIdList,
	})

	var resList []*entity.User
	for _, v := range result.UserList {
		if value, ok := userInfo.UserList[v.GetId()]; ok {
			resList = append(resList, &entity.User{
				Id:              value.GetId(),
				Name:            value.GetName(),
				Avatar:          value.GetAvatar(),
				BackgroundImage: value.GetBackgroundImage(),
				Signature:       value.GetSignature(),
			})
		}
	}

	return &relation.DouyinRelationFollowerListResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserList:   resList,
	}, err
}

func (s *Handler) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	result, err := s.clients.Relation.Client.ListRelation(ctx, &relationDomain.DoutokListRelationRequest{
		UserId:     req.UserId,
		ActionType: 2,
	})

	var friendUserList []*entity.FriendUser
	for _, v := range result.UserList {
		messageInfo, err := s.clients.Message.Client.ListMessage(ctx, &messageDomain.DoutokListMessageRequest{
			UserId:   req.UserId,
			ToUserId: v.GetId(),
		})
		if err != nil {
			continue
		}
		var msgType int64
		if len(messageInfo.Message) > 0 {
			if messageInfo.Message[len(messageInfo.Message)-1].ToUserId == req.UserId {
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
