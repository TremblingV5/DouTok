package main

import (
	"net"

	"github.com/TremblingV5/DouTok/applications/message/handler"
	"github.com/TremblingV5/DouTok/applications/message/rpc"
	"github.com/TremblingV5/DouTok/applications/message/service"
	"github.com/TremblingV5/DouTok/kitex_gen/message/messageservice"
	"github.com/cloudwego/kitex/server"

	"github.com/TremblingV5/DouTok/pkg/dlog"
)

var (
	Logger = dlog.InitLog(3)
)

func Init() {
	service.InitDb()
	service.InitHB()
	rpc.InitRPCConfig()
}

func main() {
	Init()

	// registry, err := etcd.NewEtcdRegistry([]string{
	// 	rpc.ClientConfig.Etcd.Address + ":" + rpc.ClientConfig.Etcd.Port,
	// })
	// if err != nil {
	// 	Logger.Fatal(err)
	// }

	addr, err := net.ResolveTCPAddr("tcp", rpc.ClientConfig.Server.Address+":"+rpc.ClientConfig.Server.Port)
	if err != nil {
		Logger.Fatal(err)
	}
	// fmt.Print(rpc.ClientConfig.Server.Name)
	svr := messageservice.NewServer(
		new(handler.MessageServiceImpl),
		server.WithServiceAddr(addr),
		// server.WithRegistry(registry),
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
