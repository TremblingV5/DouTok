package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/favorite/dal/db"
	"github.com/TremblingV5/DouTok/applications/favorite/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService creates a new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// FavoriteList returns a Favorite List
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*feed.Video, error) {
	FavoriteVideos, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos, err := pack.FavoriteVideos(s.ctx, FavoriteVideos, &req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
