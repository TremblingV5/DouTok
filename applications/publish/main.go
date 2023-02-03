package main

import (
	"github.com/TremblingV5/DouTok/applications/publish/handler"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
)

var (
	Logger = dlog.InitLog(3)
)

func Init() {
	service.InitDb()
	service.InitHB()
	service.InitOSS()
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

	// addr, err := net.ResolveTCPAddr("tcp", rpc.ClientConfig.Server.Address+":"+rpc.ClientConfig.Server.Port)
	// if err != nil {
	// 	Logger.Fatal(err)
	// }

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		// server.WithServiceAddr(addr),
		// server.WithRegistry(registry),
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
