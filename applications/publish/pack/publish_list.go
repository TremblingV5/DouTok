package pack

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/applications/publish/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func PackPublishListRes(list []typedef.VideoInHB, code int32, msg string, req *publish.DouyinPublishListRequest) (*publish.DouyinPublishListResponse, error) {
	res := publish.DouyinPublishListResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	newReq := user.DouyinUserRequest{
		UserId: req.UserId,
		// Token:  req.Token,
	}

	resp, _ := rpc.GetUserById(
		context.Background(), &newReq,
	)

	var video_list []*feed.Video

	for _, v := range list {
		var temp feed.Video

		temp.Title = v.GetTitle()
		temp.PlayUrl = v.GetVideoUrl()
		temp.CoverUrl = v.GetCoverUrl()

		temp.Author = resp.User
		video_list = append(video_list, &temp)
	}

	res.VideoList = video_list

	return &res, nil
}
