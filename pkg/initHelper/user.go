package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type UserClient struct {
	client userservice.Client
}

func InitUserRPCClient() *UserClient {
	config := dtviper.ConfigInit("DOUTOK_USER", "user")
	c, err := userservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &UserClient{client: c}
}

func (c *UserClient) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	resp, err := c.client.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserClient) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	resp, err := c.client.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserClient) GetUserById(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	resp, err := c.client.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
