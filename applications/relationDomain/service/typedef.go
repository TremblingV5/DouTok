package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/relation/service"
	"github.com/TremblingV5/DouTok/applications/relationDomain/dal/model"
	relationRepo "github.com/TremblingV5/DouTok/applications/relationDomain/dal/repository/relation"
	"github.com/TremblingV5/DouTok/applications/relationDomain/pack"
	"github.com/TremblingV5/DouTok/applications/relationDomain/redis/followCountRedis"
	"github.com/TremblingV5/DouTok/applications/relationDomain/redis/followListRedis"
	"github.com/TremblingV5/DouTok/applications/relationDomain/redis/followerCountRedis"
	"github.com/TremblingV5/DouTok/applications/relationDomain/redis/followerListRedis"
	entity "github.com/TremblingV5/DouTok/kitex_gen/entity"
	redishandle "github.com/TremblingV5/DouTok/pkg/redisHandle"
)

type Service struct {
	followListRedis    FollowListRedis
	followerListRedis  FollowerListRedis
	followCountRedis   FollowCountRedis
	followerCountRedis FollowerCountRedis
	relationRepo       RelationRepo
}

func New(repo *relationRepo.PersistRepository) *Service {
	redisClient := redishandle.RedisClient{
		Client: service.RedisClient,
	}

	return &Service{
		followListRedis:    followListRedis.NewClient(&redisClient),
		followerListRedis:  followerListRedis.NewClient(&redisClient),
		followCountRedis:   followCountRedis.NewClient(&redisClient),
		followerCountRedis: followerCountRedis.NewClient(&redisClient),
		relationRepo:       repo,
	}
}

type FollowListRedis interface {
	Get(ctx context.Context, userId int64) ([]int64, error)
	Set(ctx context.Context, userId int64, relations []*model.Relation) error
}

type FollowerListRedis interface {
	Get(ctx context.Context, userId int64) ([]int64, error)
	Set(ctx context.Context, userId int64, relations []*model.Relation) error
}

type FollowCountRedis interface {
	Get(ctx context.Context, userId int64) (int64, error)
	Set(ctx context.Context, userId int64, count int64) error
	Del(ctx context.Context, userId int64) error
}

type FollowerCountRedis interface {
	Get(ctx context.Context, userId int64) (int64, error)
	Set(ctx context.Context, userId int64, count int64) error
	Del(ctx context.Context, userId int64) error
}

type RelationRepo interface {
	CreateOrUpdate(relation *pack.Relation) error
	CreateList(relationList []*pack.Relation) error
}

type IService interface {
	AddRelation(ctx context.Context, userId, toUserId int64) error
	RmRelation(ctx context.Context, userId, toUserId int64) error
	ListFollowList(ctx context.Context, userId int64) ([]*entity.User, error)
	ListFollowerList(ctx context.Context, userId int64) ([]*entity.User, error)
	ListFriendList(ctx context.Context, userId int64) ([]*entity.User, error)
	GetFollowCount(ctx context.Context, userId int64) (int64, error)
	GetFollowerCount(ctx context.Context, userId int64) (int64, error)
	GetFriendCount(ctx context.Context, userId int64) (int64, error)
}
