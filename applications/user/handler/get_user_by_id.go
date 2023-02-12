package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/service"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (s *UserServiceImpl) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	user, err := service.QueryUserByIdInRDB(req.UserId)

	if err != nil {
		return service.PackUserResp(int32(misc.SearchErr.ErrCode), misc.SearchErr.ErrMsg, user)
	}

	return service.PackUserResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, user)
}
