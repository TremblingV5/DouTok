package rpc

import (
	"context"
	"errors"

	"github.com/TremblingV5/DouTok/kitex_gen/favorite"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func InitFavoriteRpc() {
	addr := ClientConfig.Etcd.Address + ":" + ClientConfig.Etcd.Port
	registry, err := etcd.NewEtcdResolver([]string{addr})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		"favorite",
		client.WithResolver(registry),
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
