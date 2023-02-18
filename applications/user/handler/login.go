package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/pack"
	"github.com/TremblingV5/DouTok/applications/user/service"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// 1. 检查参数是否非空
	if req.Username == "" || req.Password == "" {
		return pack.PackLoginResp(int32(misc.EmptyErr.ErrCode), misc.EmptyErr.ErrMsg, 0)
	}

	// 2. 检查用户名和密码是否匹配
	user_id, err, errNo := service.CheckPassword(req.Username, req.Password)
	if err != nil || user_id == 0 {
		return pack.PackLoginResp(int32(errNo.ErrCode), errNo.ErrMsg, 0)
	}

	return pack.PackLoginResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, user_id)
}
