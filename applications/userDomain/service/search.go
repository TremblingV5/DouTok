package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{
		ctx: ctx,
	}
}

func QueryUserByIdInRDB(user_id int64) (*model.User, error) {
	user, err := Do.Where(
		User.ID.Eq(uint64(user_id)),
	).First()

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *QueryUserService) QueryUserListByIdInRDB(user_id ...uint64) ([]*model.User, error) {
	userList, err := Do.Where(
		User.ID.In(user_id...),
	).Find()

	if err != nil {
		return userList, err
	}

	return userList, nil
}

func FindUserByUserName(username string) (*model.User, error) {
	res, err := Do.Where(
		User.UserName.Eq(username),
	).First()

	if err != nil {
		return &model.User{}, err
	}

	return res, nil
}
