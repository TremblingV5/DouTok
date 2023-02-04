package command

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/applications/user/dal/pack"
	"github.com/TremblingV5/DouTok/applications/user/db"
	"github.com/TremblingV5/DouTok/kitex_gen/user"

	"gorm.io/gorm"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser get user info by userID
func (s *MGetUserService) MGetUser(req *user.DouyinUserRequest, fromID int64) (*user.User, error) {
	modelUser, err := db.GetUserByID(s.ctx, req.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user, err := pack.User(s.ctx, modelUser, fromID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
