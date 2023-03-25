package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/pack"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
)

func (s *FavoriteDomainServiceImpl) AddFavorite(ctx context.Context, req *favoriteDomain.DoutokAddFavRequest) (resp *favoriteDomain.DoutokAddFavResponse, err error) {
	errNo, _ := service.NewActionFavoriteService(ctx).ActionFavorite(req.UserId, req.VideoId, true)
	return pack.PackageAddFavoriteResp(errNo)
}
