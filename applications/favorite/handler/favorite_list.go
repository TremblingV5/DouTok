package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"

	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	result, err := rpc.FavoriteDomainClient.ListFavorite(ctx, &favoriteDomain.DoutokListFavRequest{
		UserId: req.UserId,
	})
	return pack.PackageFavoriteListResponse(ctx, result, err)
}
