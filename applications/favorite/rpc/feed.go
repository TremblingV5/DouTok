package rpc

import (
	"context"

	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/kitex_gen/feed"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedservice.Client

func InitRelationRpc() {
	addr := misc.GetConfig("Etcd.Address") + ":" + misc.GetConfig("Etcd.Port")
	registry, err := etcd.NewEtcdResolver([]string{addr})
	if err != nil {
		panic(err)
	}

	c, err := feedservice.NewClient(
		"feed",
		client.WithResolver(registry),
	)

	if err != nil {
		panic(err)
	}

	feedClient = c
}

func GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	return feedClient.GetVideoById(ctx, req)
}
