package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"

	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
)

func (s *FeedServiceImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *entity.Video, err error) {
	return rpc.VideoDomainClient.GetVideoInfo(ctx, &videoDomain.DoutokGetVideoInfoRequest{
		VideoId: req.VideoId,
	})
}
