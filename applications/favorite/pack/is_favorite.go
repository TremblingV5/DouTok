package pack

import "github.com/TremblingV5/DouTok/kitex_gen/favorite"

func PackIsFavoriteResp(code int32, msg string, list map[int64]bool) (resp *favorite.DouyinIsFavoriteResponse, err error) {
	resp = &favorite.DouyinIsFavoriteResponse{
		StatusCode: code,
		StatusMsg:  msg,
		Result:     list,
	}

	return resp, nil
}
