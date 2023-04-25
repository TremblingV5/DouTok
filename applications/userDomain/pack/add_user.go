package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageAddUserResp(errno *errno.ErrNo, user_id int64) (resp *userDomain.DoutokAddUserResponse, err error) {
	return &userDomain.DoutokAddUserResponse{
		StatusCode: int32(errno.ErrCode),
		StatusMsg:  errno.ErrMsg,
		UserId:     user_id,
	}, nil
}
