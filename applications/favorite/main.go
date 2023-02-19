package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"net"

	"github.com/TremblingV5/DouTok/applications/favorite/handler"
	"github.com/TremblingV5/DouTok/applications/favorite/misc"
	"github.com/TremblingV5/DouTok/applications/favorite/rpc"
	"github.com/TremblingV5/DouTok/applications/favorite/service"
	"github.com/TremblingV5/DouTok/kitex_gen/favorite/favoriteservice"
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

	svr := favoriteservice.NewServer(
		new(handler.FavoriteServiceImpl),
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
