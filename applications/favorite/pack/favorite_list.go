package pack

import (
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"golang.org/x/net/context"
)

func PackageFavoriteListResponse(ctx context.Context, result *favoriteDomain.DoutokListFavResponse, e error) (resp *favorite.DouyinFavoriteListResponse, err error) {
	if e != nil {
		return nil, e
	}

	var videoIdList []*entity.Video
	for _, v := range result.VideoList {
		videoInfo, e := rpc.VideoDomainClient.GetVideoInfo(ctx, &videoDomain.DoutokGetVideoInfoRequest{
			VideoId: v.Id,
		})
		if e != nil {
			continue
		}
		videoIdList = append(videoIdList, videoInfo)
	}

	return &favorite.DouyinFavoriteListResponse{
		StatusCode: result.StatusCode,
		StatusMsg:  result.StatusMsg,
		VideoList:  videoIdList,
	}, nil
}
