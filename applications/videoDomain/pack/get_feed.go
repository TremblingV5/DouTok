package pack

import (
	"strconv"

	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageFeedListResp(list []typedef.VideoInHB, errNo *errno.ErrNo, userId int64) (resp *videoDomain.DoutokGetFeedResponse, err error) {
	res := videoDomain.DoutokGetFeedResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
	}

	nextTime := "9999999999"
	var videoList []*entity.Video

	for _, v := range list {
		var temp entity.Video

		temp.Id = v.GetId()
		temp.Author = &entity.User{
			Id: v.GetAuthorId(),
		}
		temp.PlayUrl = v.GetVideoUrl()
		temp.CoverUrl = v.GetCoverUrl()
		temp.Title = v.GetTitle()

		videoList = append(videoList, &temp)

		if v.GetTimestamp() < nextTime {
			nextTime = v.GetTimestamp()
		}
	}

	res.VideoList = videoList
	nextTimeInt64, _ := strconv.Atoi(nextTime)
	res.NextTime = int64(nextTimeInt64)

	return &res, nil
}
