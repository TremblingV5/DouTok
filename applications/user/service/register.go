package service

import (
	"github.com/TremblingV5/DouTok/applications/user/dal/model"
	"github.com/TremblingV5/DouTok/applications/user/misc"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/TremblingV5/DouTok/pkg/utils"
)

func WriteNewUser(username string, password string) (int64, error, *errno.ErrNo) {
	count, err := Do.Where(
		User.UserName.Eq(username),
	).Count()

	if err != nil {
		return 0, err, &misc.UserNameErr
	}

	if count > 0 {
		return 0, nil, &misc.UserNameExistedErr
	}

	user_id := utils.GetSnowFlakeId()
	salt := GenSalt()
	encrypted := PasswordEncrypt(int64(user_id), password, salt)

	if err := Do.Create(&model.User{
		ID:              uint64(user_id),
		UserName:        username,
		Password:        encrypted,
		Salt:            salt,
		Avatar:          misc.GetUserAvatar(),
		BackgroundImage: misc.GetUserAvatar(),
		Signature:       "这个人很低调",
	}); err != nil {
		return 0, err, &misc.SystemErr
	}

	return int64(user_id), nil, &misc.Success
}
