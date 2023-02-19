package main

import (
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
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
	misc.InitViperConfig()

	service.InitDb(
		misc.GetConfig("MySQL.Username"),
		misc.GetConfig("MySQL.Password"),
		misc.GetConfig("MySQL.Host"),
		misc.GetConfig("MySQL.Port"),
		misc.GetConfig("MySQL.HBase"),
	)
	service.InitHB(
		misc.GetConfig("HBase.Host"),
	)
	service.InitOSS(
		misc.GetConfig("OSS.Endpoint"),
		misc.GetConfig("OSS.Key"),
		misc.GetConfig("OSS.Secret"),
		misc.GetConfig("OSS.Bucket"),
	)
	rpc.InitPRCClient()
	utils.InitSnowFlake(1024)
}

func main() {
	Init()

	registry, err := etcd.NewEtcdRegistry([]string{
		misc.GetConfig("Etcd.Address") + ":" + misc.GetConfig("Etcd.Port"),
	})
	if err != nil {
		Logger.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", misc.GetConfig("Server.Address")+":"+misc.GetConfig("Server.Port"))
	if err != nil {
		Logger.Fatal(err)
	}

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: misc.GetConfig("Server.Name"),
			},
		),
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
