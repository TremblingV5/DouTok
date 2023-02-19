package main

import (
	"github.com/TremblingV5/DouTok/applications/favorite/handler"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/applications/favorite/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
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
		misc.GetConfig(misc.ConfigIndex_MySQLUsername),
		misc.GetConfig(misc.ConfigIndex_MySQLPassword),
		misc.GetConfig(misc.ConfigIndex_MySQLHost),
		misc.GetConfig(misc.ConfigIndex_MySQLPort),
		misc.GetConfig(misc.ConfigIndex_MySQLDb),
	)

	redisMap := map[string]int{
		misc.FavCache:    int(misc.GetConfigNum(misc.ConfigIndex_RedisFavCacheDbNum)),
		misc.FavCntCache: int(misc.GetConfigNum(misc.ConfigIndex_RedisFavCntCacheDbNum)),
	}
	service.InitRedis(
		misc.GetConfig(misc.ConfigIndex_RedisDest),
		misc.GetConfig(misc.ConfigIndex_RedisPassword),
		redisMap,
	)
	service.InitMemoryMap()

	kafkaBrokers := []string{
		misc.GetConfig("Kafka.Broker"),
	}
	service.InitKafka(kafkaBrokers)

	rpc.InitPRCClient()

	utils.InitSnowFlake(
		misc.GetConfigNum(misc.ConfigIndex_SnowFlake),
	)

	go service.UpdateFavMap()
}

func main() {
	Init()

	svr := favoriteservice.NewServer(
		new(handler.FavoriteServiceImpl),
		initHelper.InitRPCServerArgs(misc.Config)...,
	)

	if err := svr.Run(); err != nil {
		Logger.Fatal(err)
	}
}
