package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/pack"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	userInfo, err := rpc.UserDomainRPCClient.GetUserInfo(ctx, &userDomain.DoutokGetUserInfoRequest{UserId: []int64{req.UserId}})
	return pack.PackageGetUserByIdResponse(ctx, userInfo, req.UserId, err)
}
