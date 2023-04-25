package pack

import (
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
)

func PackageLoginResponse(result *userDomain.DoutokCheckUserResponse, e error) (resp *user.DouyinUserLoginResponse, err error) {
	if e != nil {
		return nil, e
	}

	return &user.DouyinUserLoginResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		UserId:     result.UserId,
	}, e
}
