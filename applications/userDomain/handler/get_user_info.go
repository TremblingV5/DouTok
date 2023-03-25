package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/misc"
	"github.com/TremblingV5/DouTok/applications/userDomain/pack"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func (s *UserDomainServiceImpl) GetUserInfo(ctx context.Context, req *userDomain.DoutokGetUserInfoRequest) (resp *userDomain.DoutokGetUserInfoResponse, err error) {
	userIdList := []uint64{}
	for _, v := range req.UserId {
		userIdList = append(userIdList, uint64(v))
	}

	userList, err := service.NewQueryUserService(ctx).QueryUserListByIdInRDB(userIdList...)
	if err != nil {
		return pack.PackageGetUserInfoResp(&misc.SystemErr, nil)
	}

	return pack.PackageGetUserInfoResp(&misc.Success, userList)
}
