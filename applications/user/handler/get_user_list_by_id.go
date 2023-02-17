package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/applications/user/pack"
	"github.com/TremblingV5/DouTok/applications/user/service"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func (s *UserServiceImpl) GetUserListByIds(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
	user_id_list := []uint64{}

	for _, v := range req.UserList {
		user_id_list = append(user_id_list, uint64(v))
	}

	user_list, err := service.QueryUserListByIdInRDB(user_id_list...)

	if err != nil {
		return pack.PackUserListResp(int32(misc.SystemErr.ErrCode), misc.SystemErr.ErrMsg, nil)
	}

	return pack.PackUserListResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, user_list)
}
