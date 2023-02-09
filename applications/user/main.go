package main

import (
	"net"

	"github.com/TremblingV5/DouTok/applications/user/handler"
	"github.com/TremblingV5/DouTok/applications/user/rpc"
	"github.com/TremblingV5/DouTok/applications/user/service"
	"github.com/TremblingV5/DouTok/kitex_gen/user/userservice"
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
	rpc.InitRPCConfig()
	utils.InitSnowFlake(1010)
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

	svr := userservice.NewServer(
		new(handler.UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
