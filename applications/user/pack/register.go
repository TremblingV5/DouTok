package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func PackageRegisterResponse(result *userDomain.DoutokAddUserResponse, e error) (resp *user.DouyinUserRegisterResponse, err error) {
	if e != nil {
		return nil, e
	}

	return &user.DouyinUserRegisterResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserId:     result.UserId,
	}, e
}
