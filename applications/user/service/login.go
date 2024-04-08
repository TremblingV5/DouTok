package service

import (
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

func CheckPassword(username string, password string) (int64, error, *errno.ErrNo) {
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
