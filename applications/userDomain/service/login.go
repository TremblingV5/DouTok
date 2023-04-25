package service

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/userDomain/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

type CheckPasswordService struct {
	ctx context.Context
}

func NewCheckPasswordService(ctx context.Context) *CheckPasswordService {
	return &CheckPasswordService{
		ctx: ctx,
	}
}

func (s *CheckPasswordService) CheckPassword(username string, password string) (int64, error, *errno.ErrNo) {
	user, err := FindUserByUserName(username)

	if err != nil {
		return 0, err, &misc.UserNameErr
	}

	encrypted := PasswordEncrypt(int64(user.ID), password, user.Salt)

	if encrypted != user.Password {
		return 0, nil, &misc.PasswordErr
	}

	return int64(user.ID), nil, &misc.Success
}
