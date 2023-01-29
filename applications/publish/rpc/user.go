package rpc

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUserRpc() {
	addr := "127.0.0.1:2379"
	registry, err := etcd.NewEtcdResolver([]string{addr})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		"user",
		client.WithResolver(registry),
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
