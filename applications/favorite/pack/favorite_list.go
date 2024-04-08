package pack

import (
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"golang.org/x/net/context"
)

func PackFavoriteListResp(code int32, msg string, list []int64) (resp *favorite.DouyinFavoriteListResponse, err error) {
	var resList []*feed.Video
	for _, i := range list {
		res, err := rpc.GetVideoById(context.Background(), &feed.VideoIdRequest{
			VideoId:  i,
			SearchId: 0,
		})
		if err != nil {
			continue
		}

		resList = append(resList, res)
	}

	return &favorite.DouyinFavoriteListResponse{
		StatusCode: code,
		StatusMsg:  msg,
		VideoList:  resList,
	}, nil
}
