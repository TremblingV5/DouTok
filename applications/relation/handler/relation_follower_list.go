package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relationDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	result, err := rpc.RelationDomainRPCClient.ListRelation(ctx, &relationDomain.DoutokListRelationRequest{
		UserId:     req.UserId,
		ActionType: 1,
	})

	var userIdList []int64
	for _, v := range result.UserList {
		userIdList = append(userIdList, v.GetId())
	}
	userInfo, err := rpc.UserDomainRPCClient.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{
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
