package service

import(
	"context"
	"github.com/TremblingV5/DouTok/applications/favorite/dal/db"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// favorite action 
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.Favorite(s.ctx, req.UserId, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.DisFavorite(s.ctx, req.UserId, req.VideoId)
	}
	return errno.ErrBind
}