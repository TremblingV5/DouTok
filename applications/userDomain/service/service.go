package service

import (
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/repository/user"
	"github.com/TremblingV5/DouTok/pkg/encrypt"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"math/rand"
)

type Service struct {
	user      user.Repository
	snowflake *utils.SnowflakeHandle
}

func New(userRepo user.Repository) *Service {
	return &Service{
		user:      userRepo,
		snowflake: utils.NewSnowflakeHandle(1),
	}
}

func (s *Service) CheckPassword(username string, password string) (int64, error) {
	user, err := s.user.LoadByUsername(username)

	if err != nil {
		return 0, err
	}

	encrypted := encrypt.Encrypt(int64(user.ID), password, user.Salt)

	if encrypted != user.Password {
		return 0, nil
	}

	return int64(user.ID), nil
}

func (s *Service) CreateNewUser(username string, password string) (int64, error) {
	exist, err := s.user.IsUserNameExisted(username)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, nil
	}

	userId := s.snowflake.GetId()
	salt := encrypt.GenSalt()
	encrypted := encrypt.Encrypt(int64(userId), password, salt)
	if err := s.user.Save(&model.User{
		ID:              uint64(userId),
		UserName:        username,
		Password:        encrypted,
		Salt:            salt,
		Avatar:          getUserAvatar(),
		BackgroundImage: getUserAvatar(),
		Signature:       "这个人很低调",
	}); err != nil {
		return 0, err
	}

	return int64(userId), nil
}

func (s *Service) LoadUserListByIds(userId ...uint64) ([]*model.User, error) {
	userList, err := s.user.LoadUserListByIds(userId...)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func getUserAvatar() string {
	link1 := "https://doutok-video.oss-cn-shanghai.aliyuncs.com/video/hear.png"
	link2 := "https://doutok-video.oss-cn-shanghai.aliyuncs.com/video/stop.png"

	r := rand.Int31n(100)
	if r%2 == 0 {
		return link1
	} else {
		return link2
	}
}
