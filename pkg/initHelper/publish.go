package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type PublishClient struct {
	client publishservice.Client
}

func InitPublishRPCClient() *PublishClient {
	config := dtviper.ConfigInit("DOUTOK_PUBLISH", "publish")
	c, err := publishservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &PublishClient{client: c}
}

func (c *PublishClient) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (*publish.DouyinPublishActionResponse, error) {
	resp, err := c.client.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *PublishClient) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (*publish.DouyinPublishListResponse, error) {
	resp, err := c.client.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
