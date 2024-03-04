package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/videoDomain/videodomainservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type VideoDomainClient struct {
	client videodomainservice.Client
}

func InitVideoDomainRPCClient() *VideoDomainClient {
	config := dtviper.ConfigInit("DOUTOK_VIDEODOMAIN", "videoDomain", nil)
	c, err := videodomainservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(config)...)
	if err != nil {
		panic(err)
	}
	return &VideoDomainClient{client: c}
}

func (c *VideoDomainClient) GetFeed(ctx context.Context, req *videoDomain.DoutokGetFeedRequest) (*videoDomain.DoutokGetFeedResponse, error) {
	resp, err := c.client.GetFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *VideoDomainClient) AddPublish(ctx context.Context, req *videoDomain.DoutokAddPublishRequest) (*videoDomain.DoutokAddPublishResponse, error) {
	resp, err := c.client.AddPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *VideoDomainClient) ListPublish(ctx context.Context, req *videoDomain.DoutokListPublishRequest) (*videoDomain.DoutokListPublishResponse, error) {
	resp, err := c.client.ListPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *VideoDomainClient) CountPublish(ctx context.Context, req *videoDomain.DoutokCountPublishRequest) (*videoDomain.DoutokCountPublishResponse, error) {
	resp, err := c.client.CountPublish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *VideoDomainClient) GetVideoInfo(ctx context.Context, req *videoDomain.DoutokGetVideoInfoRequest) (*entity.Video, error) {
	resp, err := c.client.GetVideoInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
