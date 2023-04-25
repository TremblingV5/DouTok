package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/publish/pack"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	result, err := rpc.VideoDomainClient.AddPublish(ctx, &videoDomain.DoutokAddPublishRequest{
		Data:   req.Data,
		Title:  req.Title,
		UserId: req.UserId,
		Name:   req.Name,
	})
	return pack.PackagePublishActionResponse(result, err)
}
