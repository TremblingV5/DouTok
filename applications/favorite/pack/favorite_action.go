package pack

import "github.com/TremblingV5/DouTok/kitex_gen/favorite"

func PackFavoriteActionResp(code int32, msg string) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	return &favorite.DouyinFavoriteActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}, nil
}
