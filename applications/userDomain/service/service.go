package service

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/model"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/query"
	"github.com/TremblingV5/DouTok/applications/userDomain/dal/repository/user"
	"github.com/TremblingV5/DouTok/pkg/dtdb"
	"github.com/TremblingV5/DouTok/pkg/encrypt"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"math/rand"
)

type Service struct {
	user      user.Repository
	snowflake *utils.SnowflakeHandle
	txHandle  *dtdb.TXHandle
}

func New(userRepo user.Repository) *Service {
	return &Service{
		user:      userRepo,
		snowflake: utils.NewSnowflakeHandle(1),
		txHandle: dtdb.NewTXHandle(func() dtdb.TX {
			return query.Q.Begin()
		}),
	}
}

func (s *Service) CheckPassword(ctx context.Context, username string, password string) (userId int64, err error) {
	ctx, persist := s.txHandle.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	user, err := s.user.LoadByUsername(ctx, username)

	if err != nil {
		return 0, err
	}

	encrypted := encrypt.Encrypt(int64(user.ID), password, user.Salt)

	if encrypted != user.Password {
		return 0, nil
	}

	return int64(user.ID), nil
}

func (s *Service) CreateNewUser(ctx context.Context, username string, password string) (userId int64, err error) {
	ctx, persist := s.txHandle.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	exist, err := s.user.IsUserNameExisted(ctx, username)
	if err != nil {
		return 0, err
	}
	if exist {
		return 0, errors.New("username existed")
	}

	userId = int64(s.snowflake.GetId())
	salt := encrypt.GenSalt()
	encrypted := encrypt.Encrypt(userId, password, salt)
	if err := s.user.Save(ctx, &model.User{
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

	return userId, nil
}

func (s *Service) LoadUserListByIds(ctx context.Context, userId ...uint64) (userList []*model.User, err error) {
	ctx, persist := s.txHandle.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	userList, err = s.user.LoadUserListByIds(ctx, userId...)
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
