package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/pack"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
)

func (s *FavoriteDomainServiceImpl) RmFavorite(ctx context.Context, req *favoriteDomain.DoutokRmFavRequest) (resp *favoriteDomain.DoutokRmFavResponse, err error) {
	errNo, _ := service.NewActionFavoriteService(ctx).ActionFavorite(req.UserId, req.VideoId, false)
	return pack.PackageRmFavoriteResp(errNo)
}
