package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/publish/pack"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	result, err := rpc.VideoDomainClient.ListPublish(ctx, &videoDomain.DoutokListPublishRequest{
		UserId: req.UserId,
	})
	return pack.PackagePublishListResponse(ctx, result, err)
}
