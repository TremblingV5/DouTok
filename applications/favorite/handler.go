package main

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// IsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.DouyinIsFavoriteRequest) (resp *favorite.DouyinIsFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteCount implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteCount(ctx context.Context, req *favorite.DouyinFavoriteCountRequest) (resp *favorite.DouyinFavoriteCountResponse, err error) {
	// TODO: Your code here...
	return
}
