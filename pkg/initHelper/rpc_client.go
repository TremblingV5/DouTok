package initHelper

import (
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

/*
	返回初始化RPC客户端所需要的一些配置，减少这部分代码的重复
*/
func InitRPCClientArgs(config *dtviper.Config) []client.Option {
	addr := config.Viper.GetString("Etcd.Address") + ":" + config.Viper.GetString("Etcd.Port")
	registry, err := etcd.NewEtcdResolver([]string{addr})
	if err != nil {
		panic(err)
	}

	return []client.Option{
		client.WithSuite(tracing.NewClientSuite()),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                         // mux
		client.WithRPCTimeout(30 * time.Second),             // rpc timeout
		client.WithConnectTimeout(30000 * time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),   // retry
		client.WithResolver(registry),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Viper.GetString("Server.Name")}),
	}
}
