package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/feed/dal/model"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
)

func PackVideoInfoResp(video *model.Video) (*feed.Video, error) {
	v := feed.Video{}

	v.Id = int64(video.ID)
	v.PlayUrl = video.VideoUrl
	v.CoverUrl = video.CoverUrl
	v.Title = video.Title

	id_list := []int64{int64(video.ID)}
	favCountResp, err := rpc.FavoriteCount(context.Background(), &favorite.DouyinFavoriteCountRequest{
		VideoIdList: id_list,
	})
	if err != nil {
		v.FavoriteCount = 0
	}

	comCountResp, err := rpc.CommentCount(context.Background(), &comment.DouyinCommentCountRequest{
		VideoIdList: id_list,
	})
	if err != nil {
		v.CommentCount = 0
	}

	v.FavoriteCount = favCountResp.Result[int64(video.ID)]
	v.CommentCount = comCountResp.Result[int64(video.ID)]
	v.IsFavorite = true

	return &v, nil
}
