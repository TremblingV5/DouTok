package handler

import (
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func packCheckUserResp(code int32, message string, user_id int64) (resp *userDomain.DoutokCheckUserResponse, err error) {
	return &userDomain.DoutokCheckUserResponse{
		StatusCode: code,
		StatusMsg:  message,
		UserId:     user_id,
	}, nil
}

func packGetUserInfoResp(code int32, message string, userList []*model.User) (resp *userDomain.DoutokGetUserInfoResponse, err error) {
	userListRes := make(map[int64]*entity.User)

	for _, v := range userList {
		userListRes[int64(v.ID)] = &entity.User{
			Id:              int64(v.ID),
			Name:            v.UserName,
			Avatar:          v.Avatar,
			BackgroundImage: v.BackgroundImage,
			Signature:       v.Signature,
		}
	}

	return &userDomain.DoutokGetUserInfoResponse{
		StatusCode: code,
		StatusMsg:  message,
		UserList:   userListRes,
	}, nil
}
