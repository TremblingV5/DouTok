package pack

import "github.com/TremblingV5/DouTok/kitex_gen/user"

func PackLoginResp(code int32, msg string, user_id int64) (resp *user.DouyinUserLoginResponse, err error) {
	resp = &user.DouyinUserLoginResponse{
		StatusCode: code,
		StatusMsg:  msg,
		UserId:     user_id,
	}

	return resp, nil
}

func PackRegisterResp(code int32, msg string, user_id int64) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = &user.DouyinUserRegisterResponse{
		StatusCode: code,
		StatusMsg:  msg,
		UserId:     user_id,
	}

	return resp, nil
}
