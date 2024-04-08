package rpc

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
)

var relationClient relationservice.Client

func InitRelationRpc() {

	config := dtviper.ConfigInit("DOUTOK_RELATION", "relation")

	c, err := relationservice.NewClient(
		config.Viper.GetString("Server.Name"),
		initHelper.InitRPCClientArgs(&config)...,
	)

	if err != nil {
		panic(err)
	}

	relationClient = c
}

func GetFollowCount(ctx context.Context, req *relation.DouyinRelationCountRequest) (resp *relation.DouyinRelationCountResponse, err error) {
	return relationClient.GetFollowCount(ctx, req)
}
