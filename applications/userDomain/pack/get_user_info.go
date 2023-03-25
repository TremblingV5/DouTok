package pack

import (
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageGetUserInfoResp(errNo *errno.ErrNo, userList []*model.User) (resp *userDomain.DoutokGetUserInfoResponse, err error) {
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
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
		UserList:   userListRes,
	}, nil
}
