package handler

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/applications/favorite/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	if req.ActionType == 1 {
		// 点赞
		errNo, _ := service.ActionFavorite(req.UserId, req.VideoId, true)
		return pack.PackFavoriteActionResp(int32(errNo.ErrCode), errNo.ErrMsg)
	} else if req.ActionType == 2 {
		// 取消点赞
		errNo, _ := service.ActionFavorite(req.UserId, req.VideoId, false)
		return pack.PackFavoriteActionResp(int32(errNo.ErrCode), errNo.ErrMsg)
	} else {
		return pack.PackFavoriteActionResp(int32(misc.BindingInvalidErr.ErrCode), misc.BindingInvalidErr.ErrMsg)
	}
}
