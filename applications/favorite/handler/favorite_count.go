package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/applications/favorite/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

func (s *FavoriteServiceImpl) FavoriteCount(ctx context.Context, req *favorite.DouyinFavoriteCountRequest) (resp *favorite.DouyinFavoriteCountResponse, err error) {
	if len(req.VideoIdList) == 0 {
		return pack.PackFavoriteCountResp(int32(misc.EmptyVideoIdListErr.ErrCode), misc.EmptyVideoIdListErr.ErrMsg, nil)
	}

	res, err := service.QueryFavoriteCount(req.VideoIdList)
	if err != nil {
		return pack.PackFavoriteCountResp(int32(misc.SystemErr.ErrCode), misc.SystemErr.ErrMsg, nil)
	}

	return pack.PackFavoriteCountResp(int32(misc.Success.ErrCode), misc.Success.ErrMsg, res)
}
