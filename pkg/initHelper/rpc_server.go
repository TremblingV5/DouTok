package initHelper

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

/*
	返回初始化RPC客户端所需要的一些配置，减少这部分代码的重复
*/
func InitRPCServerArgs(config *dtviper.Config) ([]server.Option, func()) {
	addr := config.Viper.GetString("Etcd.Address") + ":" + config.Viper.GetString("Etcd.Port")

	registry, err := etcd.NewEtcdRegistry([]string{addr})
	if err != nil {
		panic(err)
	}

	serverAddr, err := net.ResolveTCPAddr("tcp", config.Viper.GetString("Server.Address")+":"+config.Viper.GetString("Server.Port"))
	if err != nil {
		panic(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.Viper.GetString("Server.Name")),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%s", config.Viper.GetString("Otel.Host"), config.Viper.GetString("Otel.Port"))),
		provider.WithInsecure(),
	)

	return []server.Option{
			server.WithServiceAddr(serverAddr),
			server.WithMiddleware(middleware.CommonMiddleware),
			server.WithMiddleware(middleware.ServerMiddleware),
			server.WithRegistry(registry),
			server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
			server.WithMuxTransport(),
			server.WithSuite(tracing.NewServerSuite()),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Viper.GetString("Server.Name")}),
		}, func() {
			p.Shutdown(context.Background())
		}
}
