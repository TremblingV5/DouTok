package service

import (
	"context"
	entity "github.com/TremblingV5/DouTok/kitex_gen/entity"
)

type Service struct {
}

func New() *Service {
	return &Service{}
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
