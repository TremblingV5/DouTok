package main

import (
	"net"

	"github.com/TremblingV5/DouTok/applications/publish/handler"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/utils"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	Logger = dlog.InitLog(3)
)

func Init() {
	service.InitDb()
	service.InitHB()
	service.InitOSS()
	rpc.InitRPCConfig()
	rpc.InitPRCClient()
	utils.InitSnowFlake(1024)
}

func main() {
	Init()

	registry, err := etcd.NewEtcdRegistry([]string{
		rpc.ClientConfig.Etcd.Address + ":" + rpc.ClientConfig.Etcd.Port,
	})
	if err != nil {
		Logger.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", rpc.ClientConfig.Server.Address+":"+rpc.ClientConfig.Server.Port)
	if err != nil {
		Logger.Fatal(err)
	}

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
