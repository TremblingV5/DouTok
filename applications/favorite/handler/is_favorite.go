package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/applications/favorite/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.DouyinIsFavoriteRequest) (resp *favorite.DouyinIsFavoriteResponse, err error) {
	res, err := service.QueryIsFavorite(req.UserId, req.VideoIdList)

	if err != nil {
		return pack.PackIsFavoriteResp(int32(misc.SystemErr.ErrCode), misc.SystemErr.ErrMsg, nil)
	}

	return pack.PackIsFavoriteResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, res)
}
