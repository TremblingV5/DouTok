package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/entity"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type FeedClient struct {
	client feedservice.Client
}

func InitFeedRPCClient() *FeedClient {
	config := dtviper.ConfigInit("DOUTOK_FEED", "feed")
	c, err := feedservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &FeedClient{client: c}
}

func (c *FeedClient) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	resp, err := c.client.GetUserFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *FeedClient) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (*entity.Video, error) {
	_, err := c.client.GetVideoById(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
