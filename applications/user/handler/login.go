package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/user/pack"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"

	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	result, err := rpc.UserDomainRPCClient.CheckUser(ctx, &userDomain.DoutokCheckUserRequest{
		Username: req.Username,
		Password: req.Password,
	})

	return pack.PackageLoginResponse(result, err)
}
