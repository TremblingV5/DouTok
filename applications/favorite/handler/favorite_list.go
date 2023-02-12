package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/applications/favorite/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	// 1. 查缓存
	res, _ := service.QueryFavListInCache(req.UserId)

	// 2. 如果缓存有则直接返回
	if len(res) > 0 {
		return pack.PackFavoriteListResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, res)
	}

	// 3. 如果缓存没有则查库
	res, err = service.QueryFavListInRDB(req.UserId)

	if err != nil {
		pack.PackFavoriteListResp(int32(misc.SystemErr.ErrCode), misc.SystemErr.ErrMsg, nil)
	}

	return pack.PackFavoriteListResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, res)
}
