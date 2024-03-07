package services

import (
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/remote/trans/gonet"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"runtime"
	"time"
)

func InitRPCClientArgs(serviceName string, etcdCfg etcdConfig) []client.Option {
	registry, err := etcd.NewEtcdResolver([]string{etcdCfg.GetAddr()})
	if err != nil {
		panic(err)
	}

	options := []client.Option{
		client.WithSuite(tracing.NewClientSuite()),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithRPCTimeout(30 * time.Second),             // rpc timeout
		client.WithConnectTimeout(30000 * time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),   // retry
		client.WithResolver(registry),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	}

	if runtime.GOOS == "windows" {
		options = append(options, client.WithTransHandlerFactory(gonet.NewCliTransHandlerFactory()))
	} else {
		options = append(options, client.WithMuxConnection(1)) // mux
	}

	return options
}
