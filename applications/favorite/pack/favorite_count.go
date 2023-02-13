package pack

import "github.com/TremblingV5/DouTok/kitex_gen/favorite"

func PackFavoriteCountResp(code int32, msg string, list map[int64]int64) (resp *favorite.DouyinFavoriteCountResponse, err error) {
	resp = &favorite.DouyinFavoriteCountResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Result:     list,
	}

	return resp, nil
}
