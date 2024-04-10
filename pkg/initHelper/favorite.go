package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type FavoriteClient struct {
	client favoriteservice.Client
}

func InitFavoriteRPCClient() *FavoriteClient {
	config := dtviper.ConfigInit("DOUTOK_FAVORITE", "favorite")
	c, err := favoriteservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(&config)...)
	if err != nil {
		panic(err)
	}
	return &FavoriteClient{client: c}
}

func (c *FavoriteClient) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (*favorite.DouyinFavoriteActionResponse, error) {
	resp, err := c.client.FavoriteAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *FavoriteClient) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (*favorite.DouyinFavoriteListResponse, error) {
	resp, err := c.client.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
