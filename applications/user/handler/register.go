package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/pack"
	"github.com/TremblingV5/DouTok/applications/user/service"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// 1. 检查参数是否非空
	if req.Username == "" || req.Password == "" {
		return pack.PackRegisterResp(int32(misc.EmptyErr.ErrCode), misc.EmptyErr.ErrMsg, 0)
	}

	// 2. 写库
	user_id, _, errNo := service.WriteNewUser(req.Username, req.Password)
	return pack.PackRegisterResp(int32(errNo.ErrCode), errNo.ErrMsg, user_id)
}
