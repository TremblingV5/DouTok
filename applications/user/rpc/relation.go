package rpc

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/dtviper"

	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func InitRelationRpc() {
	config := dtviper.ConfigInit("DOUTOK_RELATION", "relation")

	addr := config.Viper.GetString("Etcd.Address") + ":" + config.Viper.GetString("Etcd.Port")
	registry, err := etcd.NewEtcdResolver([]string{addr})
	if err != nil {
		panic(err)
	}

	c, err := relationservice.NewClient(
		"relation",
		client.WithResolver(registry),
	)

	if err != nil {
		panic(err)
	}

	relationClient = c
}

func GetFollowCount(ctx context.Context, req *relation.DouyinRelationCountRequest) (resp *relation.DouyinRelationCountResponse, err error) {
	return relationClient.GetFollowCount(ctx, req)
}
