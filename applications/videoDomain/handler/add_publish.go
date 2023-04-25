package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/pack"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func (s *VideoDomainServiceImpl) AddPublish(ctx context.Context, req *videoDomain.DoutokAddPublishRequest) (resp *videoDomain.DoutokAddPublishResponse, err error) {
	if len(req.Data) == 0 || len(req.Title) == 0 {
		return pack.PackageAddPublishResp(&misc.EmptyErr)
	}

	if err := service.NewSavePublishService(ctx).SavePublish(req.UserId, req.Title, req.Data); err != nil {
		return pack.PackageAddPublishResp(&misc.SystemErr)
	}

	return pack.PackageAddPublishResp(&misc.Success)
}
