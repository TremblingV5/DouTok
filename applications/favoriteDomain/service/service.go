package service

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favoriteDomain/cache/favorite_count_cache"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/repository/favorite"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/dal/repository/favorite_count"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/redis/favorite_count_redis"
	"github.com/TremblingV5/DouTok/applications/favoriteDomain/redis/favorite_redis"
)

type Service struct {
	favorite           favorite.Repository
	favoriteCount      favorite_count.Repository   //nolint
	favoriteCountCache *favorite_count_cache.Cache //nolint
	favoriteRedis      *favorite_redis.Redis
	favoriteCountRedis *favorite_count_redis.Redis //nolint
}

type IService interface {
	CreateFavorite(ctx context.Context, userId, videoId int64) error
	RemoveFavorite(ctx context.Context, userId, videoId int64) error
	CountFavorite(ctx context.Context, videoId []int64) (map[int64]int64, error)
	ListFavorite(ctx context.Context, userId int64) ([]int64, error)
	IsFavorite(ctx context.Context, userId, videoId int64) (bool, error)
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateFavorite(ctx context.Context, userId, videoId int64) error {
	return nil
}

func (s *Service) RemoveFavorite(ctx context.Context, userId, videoId int64) error {
	return nil
}

func (s *Service) CountFavorite(ctx context.Context, videoId []int64) (map[int64]int64, error) {
	return nil, nil
}

func (s *Service) ListFavorite(ctx context.Context, userId int64) ([]int64, error) {
	return nil, nil
}

func (s *Service) IsFavorite(ctx context.Context, userId, videoId int64) (bool, error) {
	cached, err := s.favoriteRedis.Load(ctx, userId, videoId)
	if err == nil && cached {
		return true, nil
	}

	return s.favorite.Load(userId, videoId)
}
