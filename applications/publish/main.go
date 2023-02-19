package main

import (
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/publish/handler"
	"github.com/TremblingV5/DouTok/applications/publish/rpc"
	"github.com/TremblingV5/DouTok/applications/publish/service"
	"github.com/TremblingV5/DouTok/kitex_gen/publish/publishservice"
	"github.com/TremblingV5/DouTok/pkg/dlog"
	"github.com/TremblingV5/DouTok/pkg/initHelper"
	"github.com/TremblingV5/DouTok/pkg/utils"
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
	utils.InitSnowFlake(misc.GetConfigNum("Snowflake.Node"))
}

func main() {
	Init()

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		initHelper.InitRPCServerArgs(misc.Config)...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
