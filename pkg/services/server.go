package services

import (
	"context"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/remote/trans/gonet"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"runtime"
)

func InitRPCServerArgs(name string, config configStruct.BaseConfig) ([]server.Option, func()) {
	etcdAddr := config.Etcd.GetAddr()
	registry, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	serverAddr, err := net.ResolveTCPAddr("tcp", config.Base.GetAddr())
	if err != nil {
		panic(err)
	}

	var p provider.OtelProvider
	if config.Otel.Enable {
		p = provider.NewOpenTelemetryProvider(
			provider.WithServiceName(name),
			provider.WithExportEndpoint(config.Otel.GetAddr()),
			provider.WithInsecure(),
		)
	}
	options := []server.Option{
		server.WithServiceAddr(serverAddr),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithRegistry(registry),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: name}),
	}
	if runtime.GOOS == "windows" {
		options = append(options,
			server.WithTransServerFactory(gonet.NewTransServerFactory()),
			server.WithTransHandlerFactory(gonet.NewSvrTransHandlerFactory()))
	}
	return options, func() {
		if config.Otel.Enable {
			_ = p.Shutdown(context.Background())
		}
	}
}
