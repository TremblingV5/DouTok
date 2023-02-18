package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/applications/feed/service"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"strconv"
)

func PackFeedListResp(list []service.VideoInHB, code int32, msg string, user_id int64) (*feed.DouyinFeedResponse, error) {
	res := feed.DouyinFeedResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	next_time := "9999999999"
	var video_list []*feed.Video

	var video_id_list []int64
	for _, v := range list {
		video_id_list = append(video_id_list, v.GetId())
	}

	is_favorite_resp, err := rpc.IsFavorite(context.Background(), &favorite.DouyinIsFavoriteRequest{
		UserId:      user_id,
		VideoIdList: video_id_list,
	})
	if err != nil {
		return nil, nil
	}
	is_favorite := is_favorite_resp.Result

	favorite_count_resp, err := rpc.FavoriteCount(context.Background(), &favorite.DouyinFavoriteCountRequest{
		VideoIdList: video_id_list,
	})
	if err != nil {
		return nil, nil
	}
	favorite_count := favorite_count_resp.Result

	comment_count_resp, err := rpc.CommentCount(context.Background(), &comment.DouyinCommentCountRequest{
		VideoIdList: video_id_list,
	})
	if err != nil {
		return nil, nil
	}
	comment_count := comment_count_resp.Result

	for _, v := range list {
		var temp feed.Video

		temp.Id = v.GetId()
		temp.PlayUrl = v.GetVideoUrl()
		temp.CoverUrl = v.GetCoverUrl()
		temp.Title = v.GetTitle()

		temp.FavoriteCount = favorite_count[v.GetId()]
		temp.CommentCount = comment_count[v.GetId()]
		temp.IsFavorite = is_favorite[v.GetId()]

		resp, err := rpc.GetUserById(
			context.Background(), &user.DouyinUserRequest{
				UserId: v.GetAuthorId(),
			},
		)

		if err != nil {
			continue
		}

		temp.Author = resp.User

		video_list = append(video_list, &temp)

		if v.GetTimestamp() < next_time {
			next_time = v.GetTimestamp()
		}
	}

	res.VideoList = video_list
	next_time_int64, _ := strconv.Atoi(next_time)
	res.NextTime = int64(next_time_int64)

	return &res, nil
}
