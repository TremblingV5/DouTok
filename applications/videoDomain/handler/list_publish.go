package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/videoDomain/misc"
	"github.com/TremblingV5/DouTok/applications/videoDomain/pack"
	"github.com/TremblingV5/DouTok/applications/videoDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
)

func (s *VideoDomainServiceImpl) ListPublish(ctx context.Context, req *videoDomain.DoutokListPublishRequest) (resp *videoDomain.DoutokListPublishResponse, err error) {
	list, err := service.NewQueryPublishListService(ctx).QueryPublishListInHBase(req.UserId)

	if err != nil {
		return pack.PackageListPublishResp(&misc.SystemErr, nil)
	}

	return pack.PackageListPublishResp(&misc.Success, list)
}
