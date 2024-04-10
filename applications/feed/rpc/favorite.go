package rpc

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"

	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
)

var favoriteClient favoriteservice.Client

func InitFavoriteRpc() {
	config := dtviper.ConfigInit("DOUTOK_FAVORITE", "favorite")

	c, err := favoriteservice.NewClient(
		config.Viper.GetString("Server.Name"),
		initHelper.InitRPCClientArgs(config)...,
	)

	if err != nil {
		panic(err)
	}

	favoriteClient = c
}

func IsFavorite(ctx context.Context, req *favorite.DouyinIsFavoriteRequest) (resp *favorite.DouyinIsFavoriteResponse, err error) {
	resp, err = favoriteClient.IsFavorite(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}

	return resp, nil
}

func FavoriteCount(ctx context.Context, req *favorite.DouyinFavoriteCountRequest) (resp *favorite.DouyinFavoriteCountResponse, err error) {
	resp, err = favoriteClient.FavoriteCount(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errors.New(resp.StatusMsg)
	}

	return resp, nil
}
