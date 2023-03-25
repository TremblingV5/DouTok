package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/pack"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"

	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	result, err := rpc.UserDomainRPCClient.AddUser(ctx, &userDomain.DoutokAddUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	return pack.PackageRegisterResponse(result, err)
}
