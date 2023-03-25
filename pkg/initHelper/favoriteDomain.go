package initHelper

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain"
	"github.com/TremblingV5/DouTok/kitex_gen/favoriteDomain/favoritedomainservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

type FavoriteDomainClient struct {
	client favoritedomainservice.Client
}

func InitFavoriteDomainRPCClient() *FavoriteDomainClient {
	config := dtviper.ConfigInit("DOUTOK_FAVORITEDOMAIN", "favoriteDomain")
	c, err := favoritedomainservice.NewClient(config.Viper.GetString("Server.Name"), InitRPCClientArgs(config)...)
	if err != nil {
		panic(err)
	}
	return &FavoriteDomainClient{client: c}
}

func (c *FavoriteDomainClient) AddFavorite(ctx context.Context, req *favoriteDomain.DoutokAddFavRequest) (*favoriteDomain.DoutokAddFavResponse, error) {
	resp, err := c.client.AddFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *FavoriteDomainClient) RmFavorite(ctx context.Context, req *favoriteDomain.DoutokRmFavRequest) (*favoriteDomain.DoutokRmFavResponse, error) {
	resp, err := c.client.RmFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *FavoriteDomainClient) ListFavorite(ctx context.Context, req *favoriteDomain.DoutokListFavRequest) (*favoriteDomain.DoutokListFavResponse, error) {
	resp, err := c.client.ListFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *FavoriteDomainClient) IsFavorite(ctx context.Context, req *favoriteDomain.DoutokIsFavRequest) (*favoriteDomain.DoutokIsFavResponse, error) {
	resp, err := c.client.IsFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}

func (c *FavoriteDomainClient) CountFavorite(ctx context.Context, req *favoriteDomain.DoutokCountFavRequest) (*favoriteDomain.DoutokCountFavResponse, error) {
	resp, err := c.client.CountFavorite(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}
	return resp, nil
}
