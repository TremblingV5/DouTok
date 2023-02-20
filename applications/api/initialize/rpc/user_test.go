package rpc

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	mock "github.com/TremblingV5/DouTok/pkg/mock/user"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	client.EXPECT().Register(gomock.Eq(context.Background()), gomock.Eq(&user.DouyinUserRegisterRequest{Username: "integration", Password: "integration"})).Return(&user.DouyinUserRegisterResponse{
		UserId: int64(10001000),
	}, nil)

	resp, err := Register(context.Background(), client, &user.DouyinUserRegisterRequest{Username: "integration", Password: "integration"})
	assert.DeepEqual(t, int64(10001000), resp.UserId)
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	client.EXPECT().Login(gomock.Eq(context.Background()), gomock.Eq(&user.DouyinUserLoginRequest{Username: "integration", Password: "integration"})).Return(&user.DouyinUserLoginResponse{UserId: int64(10001000)}, nil)

	resp, err := Login(context.Background(), client, &user.DouyinUserLoginRequest{Username: "integration", Password: "integration"})
	assert.DeepEqual(t, int64(10001000), resp)
	assert.Nil(t, err)
}

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockClient(ctrl)
	client.EXPECT().GetUserById(gomock.Eq(context.Background()), gomock.Eq(&user.DouyinUserRequest{
		UserId: int64(10001000),
	})).Return(&user.DouyinUserResponse{User: &user.User{
		Id:              int64(10001000),
		Name:            "integration",
		FollowerCount:   int64(10),
		FollowCount:     int64(10),
		IsFollow:        true,
		Avatar:          "integration",
		BackgroundImage: "integration",
		Signature:       "integration",
		TotalFavorited:  int64(10),
		WorkCount:       int64(10),
		FavoriteCount:   int64(10),
	}}, nil)

	resp, err := GetUserById(context.Background(), client, &user.DouyinUserRequest{UserId: int64(10001000)})
	assert.DeepEqual(t, int64(10001000), resp.User.Id)
	assert.DeepEqual(t, int64(10), resp.User.FollowCount)
	assert.DeepEqual(t, int64(10), resp.User.FollowerCount)
	assert.DeepEqual(t, int64(10), resp.User.TotalFavorited)
	assert.DeepEqual(t, int64(10), resp.User.WorkCount)
	assert.DeepEqual(t, int64(10), resp.User.FavoriteCount)
	assert.DeepEqual(t, true, resp.User.IsFollow)
	assert.DeepEqual(t, "integration", resp.User.Name)
	assert.DeepEqual(t, "integration", resp.User.Avatar)
	assert.DeepEqual(t, "integration", resp.User.BackgroundImage)
	assert.DeepEqual(t, "integration", resp.User.Signature)
	assert.Nil(t, err)
}
