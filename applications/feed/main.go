package main

import (
	"github.com/TremblingV5/DouTok/applications/feed/handler"
	"github.com/TremblingV5/DouTok/applications/feed/misc"
	"github.com/TremblingV5/DouTok/applications/feed/rpc"
	"github.com/TremblingV5/DouTok/applications/feed/service"
	"github.com/TremblingV5/DouTok/kitex_gen/feed/feedservice"
	"github.com/TremblingV5/DouTok/pkg/constants"
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

	svr := feedservice.NewServer(
		new(handler.FeedServiceImpl),
		initHelper.InitRPCServerArgs(misc.Config)...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
