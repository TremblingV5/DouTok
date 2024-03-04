package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/messageDomain/messagedomainservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type MessageDomainClient struct {
	client messagedomainservice.Client
}

func InitMessageDomainRPCClient() *MessageDomainClient {
	config := dtviper.ConfigInit("DOUTOK_MESSAGEDOMAIN", "messageDomain", nil)
	c, err := messagedomainservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(config)...)
	if err != nil {
		panic(err)
	}
	return &MessageDomainClient{client: c}
}

func (c *MessageDomainClient) AddMessage(ctx context.Context, req *messageDomain.DoutokAddMessageRequest) (*messageDomain.DoutokAddMessageResponse, error) {
	resp, err := c.client.AddMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *MessageDomainClient) ListMessage(ctx context.Context, req *messageDomain.DoutokListMessageRequest) (*messageDomain.DoutokListMessageResponse, error) {
	resp, err := c.client.ListMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
