package services

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func InitRPCServerArgs(serviceName string, base baseConfig, etcdCfg etcdConfig, otelCfg otelConfig) ([]server.Option, func()) {
	etcdAddr := etcdCfg.GetAddr()
	registry, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	serverAddr, err := net.ResolveTCPAddr("tcp", base.GetAddr())
	if err != nil {
		panic(err)
	}

	var p provider.OtelProvider
	if otelCfg.IsEnable() {
		p = provider.NewOpenTelemetryProvider(
			provider.WithServiceName(serviceName),
			provider.WithExportEndpoint(otelCfg.GetAddr()),
			provider.WithInsecure(),
		)
	}

	return []server.Option{
			server.WithServiceAddr(serverAddr),
			server.WithMiddleware(middleware.CommonMiddleware),
			server.WithMiddleware(middleware.ServerMiddleware),
			server.WithRegistry(registry),
			server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
			server.WithMuxTransport(),
			server.WithSuite(tracing.NewServerSuite()),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		}, func() {
			if otelCfg.IsEnable() {
				_ = p.Shutdown(context.Background())
			}
		}
}
