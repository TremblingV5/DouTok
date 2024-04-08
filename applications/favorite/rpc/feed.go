package rpc

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"

	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
)

var feedClient feedservice.Client

func InitRelationRpc() {
	config := dtviper.ConfigInit("DOUTOK_FEED", "feed")

	c, err := feedservice.NewClient(
		config.Viper.GetString("Server.Name"),
		initHelper.InitRPCClientArgs(config)...,
	)

	if err != nil {
		panic(err)
	}

	feedClient = c
}

func GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	return feedClient.GetVideoById(ctx, req)
}
