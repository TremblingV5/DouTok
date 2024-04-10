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
	validateMap := make(map[int64]bool)
	var userIdList []int64
	validateUserIdMap := make(map[int64]bool)
	for _, v := range list {
		videoId := v.GetId()
		if _, ok := validateMap[videoId]; !ok {
			video_id_list = append(video_id_list, v.GetId())
			validateMap[videoId] = true
		}
		userId := v.GetAuthorId()
		if _, ok := validateUserIdMap[userId]; !ok {
			userIdList = append(userIdList, userId)
			validateUserIdMap[userId] = true
		}
		if v.GetTimestamp() < next_time {
			next_time = v.GetTimestamp()
		}
	}

	userInfoResp, err := rpc.GetUserListByIds(context.Background(), &user.DouyinUserListRequest{
		UserList: userIdList,
	})
	if err != nil {
		return nil, nil
	}
	userInfo := userInfoResp.UserList
	userMap := make(map[int64]*user.User)
	for _, v := range userInfo {
		userMap[v.Id] = v
	}

	isFavoriteResp, err := rpc.IsFavorite(context.Background(), &favorite.DouyinIsFavoriteRequest{
		UserId:      user_id,
		VideoIdList: video_id_list,
	})
	if err != nil {
		return nil, nil
	}
	isFavorite := isFavoriteResp.Result

	favoriteCountResp, err := rpc.FavoriteCount(context.Background(), &favorite.DouyinFavoriteCountRequest{
		VideoIdList: video_id_list,
	})
	if err != nil {
		return nil, err
	}
	favoriteCount := favoriteCountResp.Result

	commentCountResp, err := rpc.CommentCount(context.Background(), &comment.DouyinCommentCountRequest{
		VideoIdList: video_id_list,
	})
	if err != nil {
		return nil, nil
	}
	commentCount := commentCountResp.Result

	for _, v := range list {
		var temp feed.Video

		temp.Id = v.GetId()
		temp.PlayUrl = v.GetVideoUrl()
		temp.CoverUrl = v.GetCoverUrl()
		temp.Title = v.GetTitle()

		value, ok := favoriteCount[v.GetId()]
		if ok {
			temp.FavoriteCount = value
		} else {
			temp.FavoriteCount = 0
		}

		commentCnt, ok := commentCount[v.GetId()]
		if ok {
			temp.CommentCount = commentCnt
		} else {
			temp.CommentCount = 0
		}

		isFav, ok := isFavorite[v.GetId()]
		if ok {
			temp.IsFavorite = isFav
		} else {
			temp.IsFavorite = false
		}

		temp.Author = userMap[v.GetAuthorId()]

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
