package main

import (
	"github.com/TremblingV5/DouTok/applications/feed/misc"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"net"

	"github.com/TremblingV5/DouTok/applications/feed/handler"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/applications/feed/service"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
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

	redisMap := map[string]int{
		constants.FeedSendBox: int(misc.GetConfigNum("Redis.SendBox.Num")),
		constants.TimeCache:   int(misc.GetConfigNum("Redis.MarkdedTime.Num")),
	}
	service.InitRedis(
		misc.GetConfig("Redis.Dest"),
		misc.GetConfig("Redis.Password"),
		redisMap,
	)

	rpc.InitPRCClient()

	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
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

	svr := feedservice.NewServer(
		new(handler.FeedServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
