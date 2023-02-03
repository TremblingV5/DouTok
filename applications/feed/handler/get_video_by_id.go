package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/feed/service"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
)

func (s *FeedServiceImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	data, err := service.GetVideoByIdInRDB(ctx, uint64(req.VideoId))
	if err != nil {
		return &feed.Video{}, err
	}
	return service.PackVideoInfoResp(data)
}
