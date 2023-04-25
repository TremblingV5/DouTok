package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/misc"
	"github.com/TremblingV5/DouTok/applications/userDomain/pack"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func (s *UserDomainServiceImpl) CheckUser(ctx context.Context, req *userDomain.DoutokCheckUserRequest) (resp *userDomain.DoutokCheckUserResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return pack.PackageCheckUserResp(&misc.EmptyErr, 0)
	}

	userId, err, errNo := service.NewCheckPasswordService(ctx).CheckPassword(req.Username, req.Password)
	if err != nil || userId == 0 {
		return pack.PackageCheckUserResp(errNo, 0)
	}

	return pack.PackageCheckUserResp(&misc.Success, userId)
}
