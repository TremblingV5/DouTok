package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/misc"
	"github.com/TremblingV5/DouTok/applications/userDomain/pack"
	"github.com/TremblingV5/DouTok/applications/userDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func (s *UserDomainServiceImpl) AddUser(ctx context.Context, req *userDomain.DoutokAddUserRequest) (resp *userDomain.DoutokAddUserResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return pack.PackageAddUserResp(&misc.EmptyErr, 0)
	}

	userId, _, errNo := service.NewWriteNewUserService(ctx).WriteNewUser(req.Username, req.Password)
	return pack.PackageAddUserResp(errNo, userId)
}
