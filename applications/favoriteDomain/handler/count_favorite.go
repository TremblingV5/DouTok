package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/misc"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/pack"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
)

func (s *FavoriteDomainServiceImpl) CountFavorite(ctx context.Context, req *favoriteDomain.DoutokCountFavRequest) (resp *favoriteDomain.DoutokCountFavResponse, err error) {
	if len(req.UserIdList) <= 0 {
		return pack.PackageCountFavoriteResp(&misc.EmptyVideoIdListErr, nil)
	}

	res, err := service.NewQueryFavoriteCountService(ctx).QueryFavoriteCount(req.UserIdList)
	if err != nil {
		return pack.PackageCountFavoriteResp(&misc.SystemErr, nil)
	}
	return pack.PackageCountFavoriteResp(&misc.Success, res)
}
