package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageCheckUserResp(errno *errno.ErrNo, user_id int64) (resp *userDomain.DoutokCheckUserResponse, err error) {
	return &userDomain.DoutokCheckUserResponse{
		StatusCode: int32(errno.ErrCode),
		StatusMsg:  errno.ErrMsg,
		UserId:     user_id,
	}, nil
}
