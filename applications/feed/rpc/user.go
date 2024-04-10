package rpc

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"

	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
)

var userClient userservice.Client

func InitUserRpc() {
	config := dtviper.ConfigInit("DOUTOK_USER", "user")

	c, err := userservice.NewClient(
		config.Viper.GetString("Server.Name"),
		initHelper.InitRPCClientArgs(config)...,
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func GetUserById(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}

	return resp, nil
}

func GetUserListByIds(ctx context.Context, req *user.DouyinUserListRequest) (resp *user.DouyinUserListResponse, err error) {
	resp, err = userClient.GetUserListByIds(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}

	return resp, nil
}
