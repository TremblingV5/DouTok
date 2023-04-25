package handler

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	if req.ActionType == 1 {
		// 点赞
		result, _ := rpc.FavoriteDomainClient.AddFavorite(ctx, &favoriteDomain.DoutokAddFavRequest{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		return pack.PackFavoriteActionResp(result.StatusCode, result.StatusMsg)
	} else if req.ActionType == 2 {
		// 取消点赞
		result, _ := rpc.FavoriteDomainClient.RmFavorite(ctx, &favoriteDomain.DoutokRmFavRequest{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		return pack.PackFavoriteActionResp(result.StatusCode, result.StatusMsg)
	} else {
		return pack.PackFavoriteActionResp(int32(misc.BindingInvalidErr.ErrCode), misc.BindingInvalidErr.ErrMsg)
	}
}
