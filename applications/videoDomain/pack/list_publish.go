package pack

import (
	"github.com/TremblingV5/DouTok/applications/videoDomain/typedef"
	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func PackageListPublishResp(errNo *errno.ErrNo, videoList []*typedef.VideoInHB) (request *videoDomain.DoutokListPublishResponse, err error) {
	var videoListRes []*entity.Video

	for _, v := range videoList {
		videoListRes = append(videoListRes, &entity.Video{
			Id: v.GetId(),
			Author: &entity.User{
				Id: v.GetAuthorId(),
			},
			Title:    v.GetTitle(),
			PlayUrl:  v.GetVideoUrl(),
			CoverUrl: v.GetCoverUrl(),
		})
	}

	return &videoDomain.DoutokListPublishResponse{
		StatusCode: int32(errNo.ErrCode),
		StatusMsg:  errNo.ErrMsg,
		VideoList:  videoListRes,
	}, nil
}
