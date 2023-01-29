package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/applications/publish/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
)

func PackPublishActionRes(code int32, msg string) (*publish.DouyinPublishActionResponse, error) {
	var resp publish.DouyinPublishActionResponse

	resp.StatusCode = code
	resp.StatusMsg = msg

	return &resp, nil
}

func PackPublishListRes(list []typedef.VideoInHB, code int32, msg string, req *publish.DouyinPublishListRequest) (*publish.DouyinPublishListResponse, error) {
	res := publish.DouyinPublishListResponse{
		StatusCode: code,
		StatusMsg:  msg,
	}

	var video_list []*feed.Video

	for _, v := range list {
		var temp feed.Video

		temp.Title = v.Title
		temp.PlayUrl = v.VideoUrl
		temp.CoverUrl = v.CoverUrl

		newReq := user.DouyinUserRequest{
			UserId: req.UserId,
			Token:  req.Token,
		}

		resp, err := rpc.GetUserById(
			context.Background(), &newReq,
		)
		if err != nil {
			continue
		}

		temp.Author = resp.User
		video_list = append(video_list, &temp)
	}

	res.VideoList = video_list

	return &res, nil
}
