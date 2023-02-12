package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/dal/model"
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

// 查询FavRelation表，
// 筛选条件：UserId字段为给定user_id、且status为true的行
// 返回查询结果：VideoId
func FavoriteList(ctx context.Context, user_id uint64) (videolist *model.FavRelation) {
	user := new(model.FavRelation)

	DB.Where(&model.FavRelation{UserId: user_id, Status: true}, "VideoId").Find(&user)

	return user
}
