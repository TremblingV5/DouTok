package services

import (
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

func InitRPCClientArgs(base baseConfig, etcdCfg etcdConfig) []client.Option {
	registry, err := etcd.NewEtcdResolver([]string{etcdCfg.GetAddr()})
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: base.GetName()}),
	}
}
