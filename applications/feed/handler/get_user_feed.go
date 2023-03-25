package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/feed/pack"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (res *feed.DouyinFeedResponse, err error) {
	result, err := rpc.VideoDomainClient.GetFeed(ctx, &videoDomain.DoutokGetFeedRequest{
		LatestTime: req.LatestTime,
		UserId:     req.UserId,
	})
	return pack.PackageFeedListResp(ctx, result, req.UserId, err)
}
