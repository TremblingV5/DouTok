package initHelper

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/userDomain/userdomainservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type UserDomainClient struct {
	client userdomainservice.Client
}

func InitUserDomainRPCClient() *UserDomainClient {
	config := dtviper.ConfigInit("DOUTOK_USERDOMAIN", "userDomain", nil)
	c, err := userdomainservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(config)...)
	if err != nil {
		panic(err)
	}
	return &UserDomainClient{client: c}
}

func (c *UserDomainClient) AddUser(ctx context.Context, req *userDomain.DoutokAddUserRequest) (*userDomain.DoutokAddUserResponse, error) {
	resp, err := c.client.AddUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserDomainClient) CheckUser(ctx context.Context, req *userDomain.DoutokCheckUserRequest) (*userDomain.DoutokCheckUserResponse, error) {
	resp, err := c.client.CheckUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *UserDomainClient) GetUserInfo(ctx context.Context, req *userDomain.DoutokGetUserInfoRequest) (*userDomain.DoutokGetUserInfoResponse, error) {
	resp, err := c.client.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
