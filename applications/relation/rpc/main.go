package main

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/applications/relation/conf"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/kitex_gen/relation/relationservice"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"time"
)

var relationClient relationservice.Client

// Relation RPC 客户端初始化
func initRelationRpc(Config *viper.Viper) {
	EtcdAddress := fmt.Sprintf("%s:%d", Config.GetString("etcd.address"), Config.GetInt("etcd.port"))
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := Config.GetString("server.name")

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(ServiceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)

	defer p.Shutdown(context.Background())

	c, err := relationservice.NewClient(
		ServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		//client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func main() {
	v := conf.InitConfig("./config", "relation")
	initRelationRpc(v)
	resp3, err := relationClient.RelationAction(context.Background(), &relation.DouyinRelationActionRequest{
		UserId:     3,
		ToUserId:   2,
		ActionType: 1,
	})

	fmt.Println(resp3)
	resp1, err := relationClient.RelationFollowList(context.Background(), &relation.DouyinRelationFollowListRequest{
		UserId: 1,
	})

	fmt.Println(resp1)
	resp2, err := relationClient.RelationFollowerList(context.Background(), &relation.DouyinRelationFollowerListRequest{
		UserId: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp2)

}
