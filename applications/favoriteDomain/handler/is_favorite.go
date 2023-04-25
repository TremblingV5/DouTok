package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/misc"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/pack"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
)

func (s *FavoriteDomainServiceImpl) IsFavorite(ctx context.Context, req *favoriteDomain.DoutokIsFavRequest) (resp *favoriteDomain.DoutokIsFavResponse, err error) {
	res, err := service.NewQueryIsFavoriteService(ctx).QueryIsFavorite(req.UserId, req.VideoId)
	if err != nil {
		return pack.PackageIsFavoriteResp(&misc.SystemErr, nil)
	}

	return pack.PackageIsFavoriteResp(&misc.Success, res)
}
