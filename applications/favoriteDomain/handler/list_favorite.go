package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/misc"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/pack"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
)

func (s *FavoriteDomainServiceImpl) ListFavorite(ctx context.Context, req *favoriteDomain.DoutokListFavRequest) (resp *favoriteDomain.DoutokListFavResponse, err error) {
	// 1. 查缓存
	res, _ := service.QueryFavListInCache(req.UserId)

	// 2. 如果缓存有则直接返回
	if len(res) > 0 {
		return pack.PackageListFavoriteResp(&misc.Success, res)
	}

	// 3. 如果缓存没有则查库
	res, err = service.QueryFavListInRDB(req.UserId)
	// 4. 将从RDB查询到的数据读入缓存
	for _, v := range res {
		if err := service.WriteFavoriteInCache(req.UserId, v, true); err != nil {
			continue
		}
	}

	if err != nil {
		return pack.PackageListFavoriteResp(&misc.SystemErr, nil)
	}

	return pack.PackageListFavoriteResp(&misc.Success, res)
}
