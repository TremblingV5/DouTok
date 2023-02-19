package initHelper

import (
	"github.com/TremblingV5/DouTok/pkg/dtviper"
	"github.com/TremblingV5/DouTok/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

/*
	返回初始化RPC客户端所需要的一些配置，减少这部分代码的重复
*/
func InitRPCServerArgs(config *dtviper.Config) []server.Option {
	addr := config.Viper.GetString("Etcd.Address") + ":" + config.Viper.GetString("Etcd.Port")

	registry, err := etcd.NewEtcdRegistry([]string{addr})
	if err != nil {
		panic(err)
	}

	serverAddr, err := net.ResolveTCPAddr("tcp", config.Viper.GetString("Server.Address")+":"+config.Viper.GetString("Server.Port"))
	if err != nil {
		panic(err)
	}

	return []server.Option{
		server.WithServiceAddr(serverAddr),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithRegistry(registry),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Viper.GetString("Server.Name")}),
	}
}
